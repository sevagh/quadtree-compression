# quadtree-compression

### 2019-11-17 update

Experimental replacement of the native QuadTreeNode with my [k-ary-tree](https://github.com/sevagh/k-ary-tree) leads to the following results:

```
benchmark                                            old ns/op       new ns/op       delta
BenchmarkSerDeAndCompressionSmallImageLowQual-8      119796651       152582939       +27.37%
BenchmarkSerDeAndCompressionSmallImageMedQual-8      257041214       311803176       +21.30%
BenchmarkSerDeAndCompressionSmallImageHighQual-8     556795517       698409345       +25.43%
BenchmarkSerDeAndCompressionLargeImageLowQual-8      7375138657      8065578710      +9.36%
BenchmarkSerDeAndCompressionLargeImageMedQual-8      13688991867     15815462583     +15.53%
BenchmarkSerDeAndCompressionLargeImageHighQual-8     33512137526     39918996051     +19.12%

benchmark                                            old allocs     new allocs     delta
BenchmarkSerDeAndCompressionSmallImageLowQual-8      2328210        2366445        +1.64%
BenchmarkSerDeAndCompressionSmallImageMedQual-8      4693502        4693499        -0.00%
BenchmarkSerDeAndCompressionSmallImageHighQual-8     9387369        9387364        -0.00%
BenchmarkSerDeAndCompressionLargeImageLowQual-8      139293584      139293579      -0.00%
BenchmarkSerDeAndCompressionLargeImageMedQual-8      261616007      261616007      +0.00%
BenchmarkSerDeAndCompressionLargeImageHighQual-8     505739766      505739764      -0.00%

benchmark                                            old bytes       new bytes       delta
BenchmarkSerDeAndCompressionSmallImageLowQual-8      71701974        60806650        -15.20%
BenchmarkSerDeAndCompressionSmallImageMedQual-8      143925098       121205822       -15.79%
BenchmarkSerDeAndCompressionSmallImageHighQual-8     288187500       242748908       -15.77%
BenchmarkSerDeAndCompressionLargeImageLowQual-8      4331160424      3525853064      -18.59%
BenchmarkSerDeAndCompressionLargeImageMedQual-8      8287873240      6766738808      -18.35%
BenchmarkSerDeAndCompressionLargeImageHighQual-8     16334861816     13382071128     -18.08%
```

It probably makes sense - going from a cache-coherent `[4]*QuadTreeNode` representation of children to a child-sibling linkedlist saves on unallocated leaf nodes in the lowest level of the quadtree (i.e. pixel-level granularity that can no longer be split into regions), but incurs costs of creating and traversing a linked list vs. having a contiguous array of 4 children (which adds up over the millions of pixels).

### Intro

Create cool GIFs and lossily compress images with quadtrees:

![jungle-gif](./samples/jungle.gif)

### Gif creation

The above GIF was generated with the following command:

```
$ ./quadtree-compression gif \
                -delayMS 1500 \
                -quality 12 \
                ./samples/jungle.png \
                ./samples/jungle.gif
```

First, a point quadtree is built from the image containing colors per quadrant, and the average color of its 4 children (NE, NW, SE, SW), and their 4 children, etc.

By generating an image at level `n`, the quadtree is only descended to depth `n`.

Finally, images at levels `[1-quality]` are collated in a GIF to produce the demo.

This produces the effect of the image "sharpening" as the color of each quadrant is replaced with the finer granularity of its descendants.

### File compression/decompression

As a toy, there are two subcommands, `compress` and `decompress`. To compress an image and create a `.quadtree` file, the quadtree from the image is serialized to an array of uint32s, and then stored with protobuf:

```
$ ./quadtree-compression compress \
                    ./samples/jungle.png \
                    ./jungle.quadtree
$
$ du -h samples/jungle.png
11M     samples/jungle.png
$
$ du -h ./jungle.quadtree
6.8M    ./jungle.quadtree
```

`-quality` can be chosen, which, as described in the gif section, cuts the quadtree off at `depth=quality`. Low qualities create dramatically smaller quadtree files (e.g. 30K for quality=5, where the full quality is 6.8M).

Quality = 5:

![jungle-lowqual](./samples/jungle_lowqual.png)

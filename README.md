# quadtree-compression

Create cool GIFs and lossily compress images with quadtrees:
![jungle-gif](./samples/jungle.gif)

### Gif creation

The above GIF was generated with the following command:

```
$ ./quadtree-compression gif \
        -delayMS 50 -maxQuality 12 -ladder \
        ./samples/jungle.png \
        ./samples/jungle.gif
```

First, a point quadtree is built from the image containing colors per quadrant, and the average color of its 4 children (NE, NW, SE, SW), and their 4 children, etc.

By generating an image at level `n`, the quadtree is only descended to depth `n`.

Finally, images at levels `[1-quality]` are collated in a GIF to produce the demo; the `-ladder` flag then appends the reverse order to "ladder" back down in quality to create a loop.

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

# quadtree-compression

Create cool GIFs with quadtrees using only the Go standard library.

![jungle-gif](./samples/jungle.gif)


### Explanation
The above GIF was generated with the following command:

```
$ ./quadtree-compression -delayMS 1000 \
                            -quality 12 \
                            ./samples/jungle.png \
                            ./samples/jungle.gif
```

First, a point quadtree is built from the image containing colors per quadrant, and the average color of its 4 children (NE, NW, SE, SW), and their 4 children, etc.

By generating an image at level `n`, the quadtree is only descended to depth `n`.

Finally, images at levels `1-quality` are collated in a GIF to produce the demo.

This produces the effect of the image "sharpening" as the color of each quadrant is replaced with the finer granularity of its descendants.

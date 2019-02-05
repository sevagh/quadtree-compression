package main

import (
	"flag"
	"fmt"
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"os"
)

func main() {
	qualityFlag := flag.Int("quality", 10, "quadtree depth (more is better quality)")
	delayFlag := flag.Int("delayMS", 500, "frame delay in ms")

	flag.Parse()

	if flag.NArg() != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] <inpath> <outpath>\n", os.Args[0])
		os.Exit(0)
	}

	inPath := flag.Args()[0]
	outPath := flag.Args()[1]

	qt, err := BuildQuadTree(inPath)
	if err != nil {
		panic(err)
	}

	outGif := &gif.GIF{}
	for i := 1; i < *qualityFlag; i++ {
		img := qt.ToImage(i)

		pImg := image.NewPaletted(img.Bounds(), palette.Plan9)
		draw.Draw(pImg, pImg.Rect, img, img.Bounds().Min, draw.Over)

		outGif.Image = append(outGif.Image, pImg)
		outGif.Delay = append(outGif.Delay, *delayFlag/10)
	}

	f, err := os.Create(outPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	gif.EncodeAll(f, outGif)
}

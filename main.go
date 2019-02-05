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
	gifFlag := flag.Bool("gif", false, "generate gif")
	compressFlag := flag.Bool("compress", false, "compress image")
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

	if *gifFlag {
		outGif := &gif.GIF{}
		for i := 1; i < *qualityFlag; i++ {
			img := qt.ToImage(i)

			pImg := image.NewPaletted(img.Bounds(), palette.Plan9)
			draw.Draw(pImg, pImg.Rect, img, img.Bounds().Min, draw.Over)

			outGif.Image = append(outGif.Image, pImg)
			outGif.Delay = append(outGif.Delay, *delayFlag/10)
		}
		createGif(outPath, outGif)
	} else if *compressFlag {
		img := qt.ToImage(*qualityFlag)
		WriteImage(img, outPath)
	} else {
		fmt.Fprintf(os.Stderr, "Usage: -h - specify one of -gif or -compress\n")
		os.Exit(0)
	}
}

func createGif(outPath string, outGif *gif.GIF) {
	f, err := os.Create(outPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	gif.EncodeAll(f, outGif)
}

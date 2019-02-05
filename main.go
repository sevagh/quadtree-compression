package main

import (
	//"flag"
	"fmt"
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s <path-to-image> <levels of quality> <path-to-out-gif>\n", os.Args[0])
		os.Exit(1)
	}

	qt, err := BuildQuadTree(os.Args[1])
	if err != nil {
		panic(err)
	}

	levelsOfQuality, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}

	outGif := &gif.GIF{}

	for i := 1; i < levelsOfQuality; i++ {
		img, maxAchieved := qt.ToImage(i)

		pImg := image.NewPaletted(img.Bounds(), palette.Plan9)
		draw.Draw(pImg, pImg.Rect, img, img.Bounds().Min, draw.Over)

		outGif.Image = append(outGif.Image, pImg)
		outGif.Delay = append(outGif.Delay, 50)

		if maxAchieved {
			break
		}
	}

	createGif(os.Args[3], outGif)
}

func createGif(outPath string, outGif *gif.GIF) {
	f, err := os.Create(outPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	gif.EncodeAll(f, outGif)
}

package main

import (
	"fmt"
	"github.com/sevagh/quadtree-compression"
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

	qt, err := quadtree_compression.BuildQuadTree(os.Args[1])
	if err != nil {
		panic(err)
	}

	levelsOfQuality, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}

	pImages := []*image.Paletted{}

	for i := 1; i < levelsOfQuality; i++ {
		img, maxAchieved := qt.ToImage(i)

		pImg := image.NewPaletted(img.Bounds(), palette.Plan9)
		draw.Draw(pImg, pImg.Rect, img, img.Bounds().Min, draw.Over)

		pImages = append(pImages, pImg)

		if maxAchieved {
			break
		}
	}

	createGif(os.Args[3], pImages)
}

func createGif(outPath string, images []*image.Paletted) {
	f, err := os.Create(outPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	anim := gif.GIF{Delay: []int{1}, Image: images}
	gif.EncodeAll(f, &anim)
}

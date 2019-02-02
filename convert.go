package main

import (
	"image"
	colorlib "image/color"
	"image/png"
	"os"
)

func (q *QuadTree) getPixel(x, y int) colorlib.Color {
	if q.root != nil {
		return getPixel(x, y, q.root, q.width, q.height)
	}
	return colorlib.RGBA{}
}

func getPixel(x, y int, qn *quadTreeNode, xCoord, yCoord int) colorlib.Color {
	if qn.children[NW] != nil && x < xCoord/2 && y < yCoord/2 {
		return getPixel(x, y, qn.children[NW], xCoord/2, yCoord/2)
	}
	if qn.children[NE] != nil && x >= xCoord/2 && y < yCoord/2 {
		return getPixel(x-xCoord/2, y, qn.children[NE], xCoord/2, yCoord/2)
	}
	if qn.children[SW] != nil && x < xCoord/2 && y >= yCoord/2 {
		return getPixel(x, y-yCoord/2, qn.children[SW], xCoord/2, yCoord/2)
	}
	if qn.children[SE] != nil && x >= xCoord/2 && y >= yCoord/2 {
		return getPixel(x-xCoord/2, y-yCoord/2, qn.children[SE], xCoord/2, yCoord/2)
	}
	return qn.color
}

func (q *QuadTree) ConvertToImage(imgPath string) error {
	img := image.NewNRGBA(image.Rect(0, 0, q.width, q.height))

	for y := 0; y < q.height; y++ {
		for x := 0; x < q.width; x++ {
			img.Set(x, y, q.getPixel(x, y))
		}
	}

	f, err := os.Create(imgPath)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		return err
	}

	return nil
}

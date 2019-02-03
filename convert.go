package main

import (
	"image"
	colorlib "image/color"
	"image/png"
	"os"
)

func WriteImage(img image.Image, imgPath string) error {
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

func (q *QuadTree) ToImage() image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, q.Width, q.Height))

	for y := 0; y < q.Height; y++ {
		for x := 0; x < q.Width; x++ {
			img.Set(x, y, q.getPixel(x, y))
		}
	}

	return img
}

func (q *QuadTree) getPixel(x, y int) colorlib.Color {
	if q.Root != nil {
		return q.Root.getPixel(x, y, q.Width, q.Height)
	}
	return colorlib.RGBA{}
}

func (qn *QuadTreeNode) getPixel(x, y, xCoord, yCoord int) colorlib.Color {
	if qn.NW != nil && x < xCoord/2 && y < yCoord/2 {
		return qn.NW.getPixel(x, y, xCoord/2, yCoord/2)
	}
	if qn.NE != nil && x >= xCoord/2 && y < yCoord/2 {
		return qn.NE.getPixel(x-xCoord/2, y, xCoord/2, yCoord/2)
	}
	if qn.SW != nil && x < xCoord/2 && y >= yCoord/2 {
		return qn.SW.getPixel(x, y-yCoord/2, xCoord/2, yCoord/2)
	}
	if qn.SE != nil && x >= xCoord/2 && y >= yCoord/2 {
		return qn.SE.getPixel(x-xCoord/2, y-yCoord/2, xCoord/2, yCoord/2)
	}
	return qn.Color
}

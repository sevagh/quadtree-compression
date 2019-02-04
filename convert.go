package main

import (
	"image"
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

func (q *QuadTree) ToImage(level int) (image.Image, bool) {
	img := image.NewNRGBA(image.Rect(0, 0, q.Width, q.Height))

	var maxQualityAchieved bool
	var pixColor PackedRGB

	for y := 0; y < q.Height; y++ {
		for x := 0; x < q.Width; x++ {
			pixColor, maxQualityAchieved = q.getPixel(x, y, level)
			img.Set(x, y, UnpackColor(pixColor))
		}
	}

	return img, maxQualityAchieved
}

func (q *QuadTree) getPixel(x, y, level int) (PackedRGB, bool) {
	if q.Root != nil {
		return q.Root.getPixel(x, y, q.Width, q.Height, 0, level)
	}
	return 0, false
}

func (qn *QuadTreeNode) getPixel(x, y, xCoord, yCoord, level, maxLevel int) (PackedRGB, bool) {
	level += 1
	if maxLevel != -1 && level >= maxLevel {
		return qn.Color, false
	}
	if qn.NW != nil && x < xCoord/2 && y < yCoord/2 {
		return qn.NW.getPixel(x, y, xCoord/2, yCoord/2, level, maxLevel)
	}
	if qn.NE != nil && x >= xCoord/2 && y < yCoord/2 {
		return qn.NE.getPixel(x-xCoord/2, y, xCoord/2, yCoord/2, level, maxLevel)
	}
	if qn.SW != nil && x < xCoord/2 && y >= yCoord/2 {
		return qn.SW.getPixel(x, y-yCoord/2, xCoord/2, yCoord/2, level, maxLevel)
	}
	if qn.SE != nil && x >= xCoord/2 && y >= yCoord/2 {
		return qn.SE.getPixel(x-xCoord/2, y-yCoord/2, xCoord/2, yCoord/2, level, maxLevel)
	}
	return qn.Color, true
}

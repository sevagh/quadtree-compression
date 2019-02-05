package main

import (
	"image"
	_ "image/jpeg"
	png "image/png"
	"os"
)

func (q *QuadTree) ToImage(level int) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, q.Width, q.Height))

	for y := 0; y < q.Height; y++ {
		for x := 0; x < q.Width; x++ {
			img.Set(x, y, DeinterleaveZOrderRGB(q.getPixel(x, y, level)))
		}
	}

	return img
}

func LoadImage(path string) (*image.Image, error) {
	imgF, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer imgF.Close()

	imgRaw, _, err := image.Decode(imgF)
	if err != nil {
		return nil, err
	}

	return &imgRaw, nil
}

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

func (q *QuadTree) getPixel(x, y, level int) uint64 {
	if q.Root != nil {
		return q.Root.getPixel(x, y, q.Width, q.Height, 0, level)
	}
	return 0
}

func (qn *QuadTreeNode) getPixel(x, y, xCoord, yCoord, level, maxLevel int) uint64 {
	level += 1
	if maxLevel > 0 && level >= maxLevel {
		return qn.Color
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
	return qn.Color
}

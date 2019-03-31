package main

import (
	"image"
	_ "image/jpeg"
	png "image/png"
	"os"
)

func (q *QuadTree) ToImage() image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, q.Width, q.Height))

	for y := 0; y < q.Height; y++ {
		for x := 0; x < q.Width; x++ {
			img.Set(x, y, UnpackColor(q.getPixel(x, y, -1)))
		}
	}

	return img
}

func LoadImage(path string) (image.Image, error) {
	imgF, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer imgF.Close()

	imgRaw, _, err := image.Decode(imgF)
	if err != nil {
		return nil, err
	}

	return imgRaw, nil
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

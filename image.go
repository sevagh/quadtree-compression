package quadtree_compression

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

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

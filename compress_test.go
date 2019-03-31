package main

import (
	"fmt"
	"testing"
)

func TestCreateImageProgressiveQuality(t *testing.T) {
	path := "./samples/jungle.png"
	outPathFmt := "./out_%d.png"

	img, err := LoadImage(path)
	if err != nil {
		t.Fatalf("Error when loading path '%s': %+v", path, err)
	}

	qt, err := BuildQuadTree(img)
	if err != nil {
		t.Fatalf("Error when creating quadtree from image '%s': %+v", path, err)
	}

	for i := 1; i < 16; i++ {
		qt, err = qt.Compress(i)
		if err != nil {
			t.Errorf("Error when compressing quadtree: %+v", err)
		}
		img := qt.ToImage()
		outErr := WriteImage(img, fmt.Sprintf(outPathFmt, i))
		if outErr != nil {
			t.Fatalf("Error when converting quadtree to image: %+v", err)
		}
	}
}

package main

import (
	"fmt"
	"image/color"
	"testing"
)

func TestCreateImageRoundTripThroughQuadTree(t *testing.T) {
	path := "./samples/jungle.png"
	outPath := "./out.png"

	qt, err := BuildQuadTree(path)
	if err != nil {
		t.Fatalf("Error when creating quadtree from image '%s': %+v", path, err)
	}

	img, _ := qt.ToImage(-1)
	outErr := WriteImage(img, outPath)
	if outErr != nil {
		t.Fatalf("Error when converting quadtree to image: %+v", err)
	}
}

func TestCreateFakeImage(t *testing.T) {
	outPath := "./fake_out.png"

	qt := QuadTree{}

	qnNE := QuadTreeNode{}
	qnNE.Color = PackColor(color.RGBA{R: 0, G: 0, B: 0, A: 255})

	qnNW := QuadTreeNode{}
	qnNW.Color = PackColor(color.RGBA{R: 255, G: 0, B: 0, A: 255})

	qnSE := QuadTreeNode{}
	qnSE.Color = PackColor(color.RGBA{R: 0, G: 255, B: 0, A: 255})

	qnSW := QuadTreeNode{}
	qnSW.Color = PackColor(color.RGBA{R: 0, G: 0, B: 255, A: 255})

	qn := QuadTreeNode{}
	qn.NE = &qnNE
	qn.NW = &qnNW
	qn.SE = &qnSE
	qn.SW = &qnSW

	qn.Color = PackColor(color.RGBA{R: 63, G: 63, B: 63, A: 255})

	qt.Root = &qn
	qt.Height = 10
	qt.Width = 10

	img, _ := qt.ToImage(-1)
	err := WriteImage(img, outPath)
	if err != nil {
		t.Fatalf("Error creating fake image '%s': %+v", outPath, err)
	}
}

func TestCreateImageProgressiveQuality(t *testing.T) {
	path := "./samples/jungle.png"
	outPathFmt := "./out_%d.png"

	qt, err := BuildQuadTree(path)
	if err != nil {
		t.Fatalf("Error when creating quadtree from image '%s': %+v", path, err)
	}

	for i := 1; i < 16; i++ {
		img, maxAchieved := qt.ToImage(i)
		outErr := WriteImage(img, fmt.Sprintf(outPathFmt, i))
		if outErr != nil {
			t.Fatalf("Error when converting quadtree to image: %+v", err)
		}
		if maxAchieved {
			t.Logf("Already reached max size on the quadtree - won't proceed past level %d", i)
			break
		}
	}
}

package main

import (
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

	outErr := qt.ConvertToImage(outPath)
	if outErr != nil {
		t.Fatalf("Error when converting quadtree to image: %+v", err)
	}
}

func TestCreateFakeImage(t *testing.T) {
	outPath := "./fake_out.png"

	qt := QuadTree{}

	qnNE := quadTreeNode{}
	qnNE.color = color.RGBA{R: 0, G: 0, B: 0, A: 255}

	qnNW := quadTreeNode{}
	qnNW.color = color.RGBA{R: 255, G: 0, B: 0, A: 255}

	qnSE := quadTreeNode{}
	qnSE.color = color.RGBA{R: 0, G: 255, B: 0, A: 255}

	qnSW := quadTreeNode{}
	qnSW.color = color.RGBA{R: 0, G: 0, B: 255, A: 255}

	qn := quadTreeNode{}
	qn.children[NE] = &qnNE
	qn.children[NW] = &qnNW
	qn.children[SE] = &qnSE
	qn.children[SW] = &qnSW

	qn.color = color.RGBA{R: 63, G: 63, B: 63, A: 255}

	qt.root = &qn
	qt.height = 10
	qt.width = 10

	err := qt.ConvertToImage(outPath)
	if err != nil {
		t.Fatalf("Error creating fake image '%s': %+v", outPath, err)
	}
}

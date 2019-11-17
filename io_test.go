package main

import (
	"image/color"
	"testing"

	"github.com/sevagh/quadtree-compression/k-ary-tree"
)

func TestLoadPngImage(t *testing.T) {
	path := "./samples/jungle.png"
	helper(path, t)
}

func TestLoadJpegImage(t *testing.T) {
	path := "./samples/jungle.jpg"
	helper(path, t)
}

func helper(path string, t *testing.T) {
	img, err := LoadImage(path)
	if err != nil {
		t.Fatalf("Error when loading path '%s': %+v", path, err)
	}

	if img == nil {
		t.Errorf("Expected valid object, got nil")
	}
}

func TestCreateImageRoundTripThroughQuadTree(t *testing.T) {
	path := "./samples/jungle.png"
	outPath := "./out.png"

	img, err := LoadImage(path)
	if err != nil {
		t.Fatalf("Error when loading path '%s': %+v", path, err)
	}

	qt, err := BuildQuadTree(img)
	if err != nil {
		t.Fatalf("Error when creating quadtree from image '%s': %+v", path, err)
	}

	img = qt.ToImage()
	outErr := WriteImage(img, outPath)
	if outErr != nil {
		t.Fatalf("Error when converting quadtree to image: %+v", err)
	}
}

func TestCreateFakeImage(t *testing.T) {
	outPath := "./fake_out.png"

	qt := QuadTree{}

	qnNE := karytree.NewNode(PackColor(color.RGBA{R: 0, G: 0, B: 0, A: 255}))
	qnNW := karytree.NewNode(PackColor(color.RGBA{R: 255, G: 0, B: 0, A: 255}))
	qnSE := karytree.NewNode(PackColor(color.RGBA{R: 0, G: 255, B: 0, A: 255}))
	qnSW := karytree.NewNode(PackColor(color.RGBA{R: 0, G: 0, B: 255, A: 255}))

	qn := karytree.NewNode(PackColor(color.RGBA{R: 63, G: 63, B: 63, A: 255}))

	qn.SetNthChild(NE, &qnNE)
	qn.SetNthChild(NW, &qnNW)
	qn.SetNthChild(SE, &qnSE)
	qn.SetNthChild(SW, &qnSW)

	qt.Root = &qn
	qt.Height = 10
	qt.Width = 10

	img := qt.ToImage()
	err := WriteImage(img, outPath)
	if err != nil {
		t.Fatalf("Error creating fake image '%s': %+v", outPath, err)
	}
}

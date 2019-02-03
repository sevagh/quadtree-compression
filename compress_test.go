package main

import (
	"image/color"
	"testing"
)

func TestCompressModifiesQuadTreeObject(t *testing.T) {
	qt := QuadTree{}

	qnNE := quadTreeNode{}
	qnNE.color = color.RGBA{R: 255, G: 0, B: 0, A: 0}

	qnNW := quadTreeNode{}
	qnNW.color = color.RGBA{R: 255, G: 0, B: 0, A: 0}

	qnSE := quadTreeNode{}
	qnSE.color = color.RGBA{R: 255, G: 0, B: 0, A: 0}

	qnSW := quadTreeNode{}
	qnSW.color = color.RGBA{R: 255, G: 0, B: 0, A: 0}

	qn := quadTreeNode{}
	qn.children[NE] = &qnNE
	qn.children[NW] = &qnNW
	qn.children[SE] = &qnSE
	qn.children[SW] = &qnSW

	qn.color = color.RGBA{R: 255, G: 0, B: 0, A: 0}

	qt.root = &qn
	qt.height = 4
	qt.width = 4

	nonNilCount := 0
	for i := 0; i < 4; i++ {
		if qt.root.children[i] != nil {
			nonNilCount += 1
		}
	}

	if nonNilCount != 4 {
		t.Errorf("Problem manually creating quadtree")
	}

	qt.Compress(1000)

	nonNilCount = 0
	for i := 0; i < 4; i++ {
		if qt.root.children[i] != nil {
			nonNilCount += 1
		}
	}

	if nonNilCount > 0 {
		t.Errorf("Expected compression step to eliminate children")
	}

	expectedColor := color.RGBA{R: 255, G: 0, B: 0, A: 0}
	if qt.root.color != expectedColor {
		t.Errorf("Expected compressed color to be R255")
	}
}

func TestCompressQuadTree(t *testing.T) {
	path := "./samples/jungle.jpg"
	regularOut := "./normal_out.png"
	compressedPath := "./compressed_out.png"

	qt, err := BuildQuadTree(path)
	if err != nil {
		t.Fatalf("Error when creating quadtree from image '%s': %+v", path, err)
	}

	if qt == nil {
		t.Errorf("Expected a non-nil quad tree")
	}

	err = qt.ConvertToImage(regularOut)
	if err != nil {
		t.Fatalf("Error when outputting initial tree to '%s': '%+v", regularOut, err)
	}

	qt.Compress(10000)

	err = qt.ConvertToImage(compressedPath)
	if err != nil {
		t.Fatalf("Error when outputting compressed tree to '%s': '%+v", compressedPath, err)
	}
}

package main

import (
	"image/color"
	"testing"
)

func TestPruneModifiesQuadTreeObject(t *testing.T) {
	qt := QuadTree{}

	qnNE := QuadTreeNode{}
	qnNE.Color = color.RGBA{R: 255, G: 0, B: 0, A: 0}

	qnNW := QuadTreeNode{}
	qnNW.Color = color.RGBA{R: 255, G: 0, B: 0, A: 0}

	qnSE := QuadTreeNode{}
	qnSE.Color = color.RGBA{R: 255, G: 0, B: 0, A: 0}

	qnSW := QuadTreeNode{}
	qnSW.Color = color.RGBA{R: 255, G: 0, B: 0, A: 0}

	qn := QuadTreeNode{}
	qn.NE = &qnNE
	qn.NW = &qnNW
	qn.SE = &qnSE
	qn.SW = &qnSW

	qn.Color = color.RGBA{R: 255, G: 0, B: 0, A: 0}

	qt.Root = &qn
	qt.Height = 4
	qt.Width = 4

	nonNilCount := 0
	if qt.Root.NE != nil {
		nonNilCount += 1
	}
	if qt.Root.NW != nil {
		nonNilCount += 1
	}
	if qt.Root.SE != nil {
		nonNilCount += 1
	}
	if qt.Root.SW != nil {
		nonNilCount += 1
	}

	if nonNilCount != 4 {
		t.Errorf("Problem manually creating quadtree")
	}

	qt.Prune()

	nonNilCount = 0

	if qt.Root.NE != nil {
		nonNilCount += 1
	}
	if qt.Root.NW != nil {
		nonNilCount += 1
	}
	if qt.Root.SE != nil {
		nonNilCount += 1
	}
	if qt.Root.SW != nil {
		nonNilCount += 1
	}

	if nonNilCount > 0 {
		t.Errorf("Expected prune step to eliminate children")
	}

	expectedColor := color.RGBA{R: 255, G: 0, B: 0, A: 0}
	if qt.Root.Color != expectedColor {
		t.Errorf("Expected prune color to be R255")
	}
}

func TestPruneQuadTree(t *testing.T) {
	path := "./samples/jungle.png"
	regularOut := "./normal_out.png"
	prunedPath := "./pruned_out.png"

	qt, err := BuildQuadTree(path)
	if err != nil {
		t.Fatalf("Error when creating quadtree from image '%s': %+v", path, err)
	}

	if qt == nil {
		t.Errorf("Expected a non-nil quad tree")
	}

	err = WriteImage(qt.ToImage(), regularOut)
	if err != nil {
		t.Fatalf("Error when outputting initial tree to '%s': '%+v", regularOut, err)
	}

	qt.Prune()

	err = WriteImage(qt.ToImage(), prunedPath)
	if err != nil {
		t.Fatalf("Error when outputting pruneed tree to '%s': '%+v", prunedPath, err)
	}
}

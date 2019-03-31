package main

import (
	"testing"
)

func TestCreateQuadTreeFromPng(t *testing.T) {
	path := "./samples/jungle.png"

	img, err := LoadImage(path)
	if err != nil {
		t.Fatalf("Error when loading image path '%s': %+v", path, err)
	}

	qt, err := BuildQuadTree(img)
	if err != nil {
		t.Fatalf("Error when creating quadtree from image '%s': %+v", path, err)
	}

	if qt == nil {
		t.Errorf("Expected a non-nil quad tree")
	}
}

package main

import (
	"testing"
)

func TestCreateQuadTreeFromPng(t *testing.T) {
	path := "./samples/jungle.png"

	qt, err := BuildQuadTree(path)
	if err != nil {
		t.Fatalf("Error when creating quadtree from image '%s': %+v", path, err)
	}

	if qt == nil {
		t.Errorf("Expected a non-nil quad tree")
	}
}

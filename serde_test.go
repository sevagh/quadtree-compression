package main

import (
	"testing"
)

func TestSerDeQuadTree(t *testing.T) {
	path := "./samples/jungle.png"
	outPathQuadTree := "./out.quadtree"

	qt, err := BuildQuadTree(path)
	if err != nil {
		t.Fatalf("Error when creating quadtree from image '%s': %+v", path, err)
	}

	err = qt.SerializeToFile(outPathQuadTree)
	if err != nil {
		t.Fatalf("Error when serializing quadtree to file '%s': %+v", outPathQuadTree, err)
	}

	qt2, err := LoadQuadTreeFromFile(outPathQuadTree)
	if err != nil {
		t.Fatalf("Error when deserializing quadtree from file '%s': %+v", outPathQuadTree, err)
	}

	outPathImage := "./out_serde.png"
	err = WriteImage(qt2.ToImage(), outPathImage)
	if err != nil {
		t.Fatalf("Error when converting quadtree to image: %+v", err)
	}
}

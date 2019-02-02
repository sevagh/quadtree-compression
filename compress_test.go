package main

import (
	"testing"
)

func TestCompressQuadTree(t *testing.T) {
	path := "./samples/jungle.png"
	compressedPath := "./compressed_out.png"

	qt, err := BuildQuadTree(path)
	if err != nil {
		t.Fatalf("Error when creating quadtree from image '%s': %+v", path, err)
	}

	if qt == nil {
		t.Errorf("Expected a non-nil quad tree")
	}

	qt.Compress(0.5)

	err = qt.ConvertToImage(compressedPath)
	if err != nil {
		t.Fatalf("Error when outputting compressed tree to '%s': '%+v", compressedPath, err)
	}
}

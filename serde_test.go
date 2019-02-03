package main

import (
	"reflect"
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

	qtSlice := qt.TreeToSlice()
	qt2Slice := qt2.TreeToSlice()

	t.Logf("leaf count: %+v\n", qt.Leaves())
	t.Logf("array count: %+v\n", len(qtSlice))
	t.Logf("leaf2 count: %+v\n", qt2.Leaves())
	t.Logf("array2 count: %+v\n", len(qt2Slice))

	if !reflect.DeepEqual(qtSlice, qt2Slice) {
		t.Fatal("Error when serializing quadtree to array and back")
	}

	outPathImage := "./out_serde.png"
	err = WriteImage(qt2.ToImage(), outPathImage)
	if err != nil {
		t.Fatalf("Error when converting quadtree to image: %+v", err)
	}
}

package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSerDeQuadTree(t *testing.T) {
	path := "./samples/jungle.png"
	outPathQuadTree := "./out.quadtree"

	img, err := LoadImage(path)
	if err != nil {
		t.Fatalf("Error when loading image from path '%s': %+v", path, err)
	}

	qt, err := BuildQuadTree(img)
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
	t.Logf("WxH count: %+v\n", qt.Width*qt.Height)
	t.Logf("array count: %+v\n", len(qtSlice))
	t.Logf("leaf2 count: %+v\n", qt2.Leaves())
	t.Logf("array2 count: %+v\n", len(qt2Slice))

	if !reflect.DeepEqual(qtSlice, qt2Slice) {
		t.Fatal("Error when serializing quadtree to array and back")
	}

	outPathImage := "./out_serde.png"
	img = qt2.ToImage()
	err = WriteImage(img, outPathImage)
	if err != nil {
		t.Fatalf("Error when converting quadtree to image: %+v", err)
	}
}

func TestSerDeCompressedWithLevelsQuadTree(t *testing.T) {
	path := "./samples/jungle.png"
	outPathQuadTreeFmt := "./out_%d.quadtree"
	outPathImageFmt := "./out_serde_%d.png"

	img, err := LoadImage(path)
	if err != nil {
		t.Fatalf("Error when loading image from path '%s': %+v", path, err)
	}

	qt, err := BuildQuadTree(img)
	if err != nil {
		t.Fatalf("Error when creating quadtree from image '%s': %+v", path, err)
	}

	for i := 16; i >= 1; i-- {
		quadTreeOutFile := fmt.Sprintf(outPathQuadTreeFmt, i)
		quadTreeImageOutFile := fmt.Sprintf(outPathImageFmt, i)

		qt, _ = qt.Compress(i)

		err = qt.SerializeToFile(quadTreeOutFile)
		if err != nil {
			t.Fatalf("Error when serializing quadtree to file '%s': %+v", quadTreeOutFile, err)
		}

		qt2, err := LoadQuadTreeFromFile(quadTreeOutFile)
		if err != nil {
			t.Fatalf("Error when deserializing quadtree from file '%s': %+v", quadTreeOutFile, err)
		}

		qtSlice := qt.TreeToSlice()
		qt2Slice := qt2.TreeToSlice()

		if !reflect.DeepEqual(qtSlice, qt2Slice) {
			t.Fatal("Error when serializing quadtree to array and back")
		}

		img := qt2.ToImage()
		err = WriteImage(img, quadTreeImageOutFile)
		if err != nil {
			t.Fatalf("Error when converting quadtree to image: %+v", err)
		}
	}
}

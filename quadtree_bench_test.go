package main

import (
	"fmt"
	"reflect"
	"testing"
)

func BenchmarkSerDeAndCompression(b *testing.B) {
	path := "./samples/jungle.png"
	outPathQuadTreeFmt := "./out_%d.quadtree"
	outPathImageFmt := "./out_serde_%d.png"

	img, err := LoadImage(path)
	if err != nil {
		b.Fatalf("Error when loading image from path '%s': %+v", path, err)
	}

	qt, err := BuildQuadTree(img)
	if err != nil {
		b.Fatalf("Error when creating quadtree from image '%s': %+v", path, err)
	}

	for j := 0; j < b.N; j++ {
		for i := 16; i >= 1; i-- {
			quadTreeOutFile := fmt.Sprintf(outPathQuadTreeFmt, i)
			quadTreeImageOutFile := fmt.Sprintf(outPathImageFmt, i)

			qt, _ = qt.Compress(i)

			err = qt.SerializeToFile(quadTreeOutFile)
			if err != nil {
				b.Fatalf("Error when serializing quadtree to file '%s': %+v", quadTreeOutFile, err)
			}

			qt2, err := LoadQuadTreeFromFile(quadTreeOutFile)
			if err != nil {
				b.Fatalf("Error when deserializing quadtree from file '%s': %+v", quadTreeOutFile, err)
			}

			qtSlice := qt.TreeToSlice()
			qt2Slice := qt2.TreeToSlice()

			// don't time comparison
			b.StopTimer()
			if !reflect.DeepEqual(qtSlice, qt2Slice) {
				b.Fatal("Error when serializing quadtree to array and back")
			}
			b.StartTimer()

			img := qt2.ToImage()
			err = WriteImage(img, quadTreeImageOutFile)
			if err != nil {
				b.Fatalf("Error when converting quadtree to image: %+v", err)
			}
		}
	}
}

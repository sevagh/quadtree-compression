package main

import (
	"testing"
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

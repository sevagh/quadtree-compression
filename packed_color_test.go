package main

import (
	"image/color"
	"testing"
)

func TestPackColor(t *testing.T) {
	myColors := []color.RGBA{
		color.RGBA{R: 137, G: 153, B: 83, A: 0xFF},
		color.RGBA{R: 137, G: 153, B: 83, A: 0xFF},
		color.RGBA{R: 194, G: 0, B: 120, A: 0xFF},
		color.RGBA{R: 203, G: 65, B: 107, A: 0xFF},
	}

	for _, myColor := range myColors {
		myColor_ := PackColor(myColor)

		origColor := UnpackColor(myColor_)

		R, G, B, A := origColor.RGBA()

		if uint8(R) != myColor.R || uint8(G) != myColor.G || uint8(B) != myColor.B || uint8(A) != myColor.A {
			t.Errorf("Expected color to be identical after pack and unpack: %+v - %+v", myColor, origColor)
		}
	}
}

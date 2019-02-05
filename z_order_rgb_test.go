package main

import (
	"image/color"
	"testing"
)

func TestZOrderColor(t *testing.T) {
	myColor := color.RGBA{R: 137, G: 153, B: 83, A: 13}

	zorder := InterleaveColor(myColor)

	origColor := DeinterleaveZOrderRGB(zorder)

	R, G, B, A := origColor.RGBA()

	if uint8(R) != myColor.R || uint8(G) != myColor.G || uint8(B) != myColor.B || uint8(A) != myColor.A {
		t.Errorf("Expected color to be identical after pack and unpack: %+v - %+v", myColor, origColor)
	}
}

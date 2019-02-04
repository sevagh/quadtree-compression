package quadtree_compression

import (
	"image/color"
	"testing"
)

func TestPackAndUnpackColor(t *testing.T) {
	myColor := color.RGBA{R: 137, G: 153, B: 83, A: 13}

	packed := PackColor(myColor)

	unpacked := UnpackColor(packed)

	R, G, B, A := unpacked.RGBA()

	if uint8(R) != myColor.R || uint8(G) != myColor.G || uint8(B) != myColor.B || uint8(A) != myColor.A {
		t.Errorf("Expected color to be identical after pack and unpack: %+v - %+v", myColor, unpacked)
	}
}

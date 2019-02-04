package quadtree_compression

import (
	"image/color"
)

type PackedRGB uint64

func PackColor(c color.Color) PackedRGB {
	R, G, B, A := c.RGBA()

	var packed uint64

	packed += (uint64(R) & 0x000000000000FFFF) |
		(uint64(G) << 16 & 0x00000000FFFF0000) |
		(uint64(B) << 32 & 0x0000FFFF00000000) |
		(uint64(A) << 48 & 0xFFFF000000000000)

	return PackedRGB(packed)
}

func UnpackColor(p PackedRGB) color.Color {
	A := uint16((p & 0xFFFF000000000000) >> 48)
	B := uint16((p & 0x0000FFFF00000000) >> 32)
	G := uint16((p & 0x00000000FFFF0000) >> 16)
	R := uint16((p & 0x000000000000FFFF))

	return color.RGBA64{R: R, G: G, B: B, A: A}
}

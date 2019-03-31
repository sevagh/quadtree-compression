package main

import (
	"image/color"
)

var SCALE_16_TO_8 uint32 = uint32(^uint16(0)) / uint32(^uint8(0))

func PackColor(c color.Color) uint32 {
	R, G, B, _ := c.RGBA()

	r := uint8(uint32(R) / SCALE_16_TO_8)
	g := uint8(uint32(G) / SCALE_16_TO_8)
	b := uint8(uint32(B) / SCALE_16_TO_8)

	var packed uint32

	packed += (uint32(r) & 0x000000FF) |
		(uint32(g) << 8 & 0x0000FF00) |
		(uint32(b) << 16 & 0x00FF0000)

	return uint32(packed)
}

func UnpackColor(p uint32) color.Color {
	b := uint8((p & 0x00FF0000) >> 16)
	g := uint8((p & 0x0000FF00) >> 8)
	r := uint8((p & 0x000000FF))

	R := uint16(uint32(r) * SCALE_16_TO_8)
	G := uint16(uint32(g) * SCALE_16_TO_8)
	B := uint16(uint32(b) * SCALE_16_TO_8)
	A := uint16(0xFFFF)

	return color.RGBA64{R: R, G: G, B: B, A: A}
}

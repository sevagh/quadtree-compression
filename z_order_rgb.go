package main

import (
	"image/color"
)

func InterleaveColor(c color.Color) uint64 {
	R, G, B, A := c.RGBA()

	var interleaved uint64

	interleaved += uint64(R & 0xF)
	interleaved += uint64(G&0xF) << 4
	interleaved += uint64(B&0xF) << 8
	interleaved += uint64(A&0xF) << 12

	interleaved += uint64(R&0xF0) << 12
	interleaved += uint64(G&0xF0) << 16
	interleaved += uint64(B&0xF0) << 20
	interleaved += uint64(A&0xF0) << 24

	interleaved += uint64(R&0xF00) << 24
	interleaved += uint64(G&0xF00) << 28
	interleaved += uint64(B&0xF00) << 32
	interleaved += uint64(A&0xF00) << 36

	interleaved += uint64(R&0xF000) << 36
	interleaved += uint64(G&0xF000) << 40
	interleaved += uint64(B&0xF000) << 44
	interleaved += uint64(A&0xF000) << 48

	return interleaved
}

func DeinterleaveZOrderRGB(z uint64) color.Color {
	var R uint16
	var G uint16
	var B uint16
	var A uint16

	A += uint16((z & 0xF000000000000000) >> 48)
	B += uint16((z & 0x0F00000000000000) >> 44)
	G += uint16((z & 0x00F0000000000000) >> 40)
	R += uint16((z & 0x000F000000000000) >> 36)

	A += uint16((z & 0xF00000000000) >> 36)
	B += uint16((z & 0x0F0000000000) >> 32)
	G += uint16((z & 0x00F000000000) >> 28)
	R += uint16((z & 0x000F00000000) >> 24)

	A += uint16((z & 0xF0000000) >> 24)
	B += uint16((z & 0x0F000000) >> 20)
	G += uint16((z & 0x00F00000) >> 16)
	R += uint16((z & 0x000F0000) >> 12)

	A += uint16((z & 0xF000) >> 12)
	B += uint16((z & 0x0F00) >> 8)
	G += uint16((z & 0x00F0) >> 4)
	R += uint16(z & 0x000F)

	return color.RGBA64{R: R, G: G, B: B, A: A}
}

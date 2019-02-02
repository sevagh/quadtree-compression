package main

import (
	"image/color"
	"math"
)

// color formulas taken from: http://www.easyrgb.com/en/math.php
// and https://en.wikipedia.org/wiki/Color_difference#CIE76

type lab struct {
	l float64
	a float64
	b float64
}

type xyz struct {
	x float64
	y float64
	z float64
}

//  \Delta E_{ab}^{*}\approx 2.3 corresponds to a JND (just noticeable difference).[6]
func CIE76(color1, color2 color.Color) float64 {
	lab1 := rgbToLab(color1)
	lab2 := rgbToLab(color2)

	deltaE := math.Sqrt(
		math.Pow(lab2.l-lab1.l, 2) +
			math.Pow(lab2.a-lab1.a, 2) +
			math.Pow(lab2.b-lab1.b, 2))

	return deltaE
}

func rgbToLab(c color.Color) lab {
	myXyz := rgbToXyz(c)
	return myXyz.toLab()
}

func (myXyz *xyz) toLab() lab {
	X := myXyz.x / 94.811
	Y := myXyz.y / 100.000
	Z := myXyz.z / 107.304

	for _, x := range []*float64{&X, &Y, &Z} {
		if *x > 0.008856 {
			*x = math.Pow(*x, 1.0/3.0)
		} else {
			*x = 7.787**x + 16.0/116.0
		}
	}

	return lab{
		l: 116.0*Y - 16.0,
		a: 500 * (X - Y),
		b: 200 * (Y - Z),
	}
}

func rgbToXyz(c color.Color) xyz {
	R_, G_, B_, _ := c.RGBA()
	R := float64(R_) / 255
	G := float64(G_) / 255
	B := float64(B_) / 255

	for _, x := range []*float64{&R, &G, &B} {
		if *x > 0.04045 {
			*x = math.Pow(((*x + 0.055) / 1.055), 2.4)
		} else {
			*x /= 12.92
		}
		*x *= 100.0
	}

	return xyz{
		x: R*0.4124 + G*0.3576 + B*0.1805,
		y: R*0.2126 + G*0.7152 + B*0.0722,
		z: R*0.0193 + G*0.1192 + B*0.9505,
	}
}

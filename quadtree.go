package main

import (
	"image"
	colorlib "image/color"
)

const (
	NE = iota
	NW
	SW
	SE
)

type quadTreeNode struct {
	children [4]*quadTreeNode //indexed with the cardinal direction iota
	color    colorlib.Color
}

type QuadTree struct {
	root   *quadTreeNode
	height int
	width  int
}

func newQuadTreeNode(color colorlib.Color) quadTreeNode {
	return quadTreeNode{color: color, children: [4]*quadTreeNode{nil, nil, nil, nil}}
}

func BuildQuadTree(imageSource string) (*QuadTree, error) {
	img, err := LoadImage(imageSource)
	if err != nil {
		return nil, err
	}

	qt := QuadTree{}

	qt.width = (*img).Bounds().Max.X - 1 - (*img).Bounds().Min.X
	qt.height = (*img).Bounds().Max.Y - 1 - (*img).Bounds().Min.Y

	qt.root = buildQuadTree(img, (*img).Bounds().Min.X, (*img).Bounds().Min.Y, qt.width, qt.height)

	return &qt, nil
}

func buildQuadTree(img *image.Image, x, y, w, h int) *quadTreeNode {
	if w == 1 && h == 1 {
		qn := newQuadTreeNode((*img).At(x, y))
		return &qn
	}

	qn := quadTreeNode{}

	qn.children[NW] = buildQuadTree(img, x, y, w/2, h/2)
	qn.children[NE] = buildQuadTree(img, x+(w/2), y, w/2, h/2)
	qn.children[SW] = buildQuadTree(img, x, y+(h/2), w/2, h/2)
	qn.children[SE] = buildQuadTree(img, x+(w/2), y+(h/2), w/2, h/2)

	if qn.children[NE] != nil && qn.children[NW] != nil && qn.children[SW] != nil && qn.children[SE] != nil {
		var red uint32
		var green uint32
		var blue uint32
		var alpha uint32

		for i := 0; i < 4; i++ {
			red_, green_, blue_, alpha_ := qn.children[i].color.RGBA()
			red += red_
			green += green_
			blue += blue_
			alpha += alpha_
		}

		red /= 4
		green /= 4
		blue /= 4
		alpha /= 4

		qn.color = colorlib.RGBA64{R: uint16(red), G: uint16(green), B: uint16(blue), A: uint16(alpha)}
	}

	return &qn
}

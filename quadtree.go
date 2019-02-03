package main

import (
	//"fmt"
	"image"
	colorlib "image/color"
)

type QuadTreeNode struct {
	NW    *QuadTreeNode
	NE    *QuadTreeNode
	SE    *QuadTreeNode
	SW    *QuadTreeNode
	Color colorlib.Color
}

type QuadTree struct {
	Root   *QuadTreeNode
	Height int
	Width  int
}

func NewQuadTreeNode(color colorlib.Color) QuadTreeNode {
	return QuadTreeNode{Color: color}
}

func BuildQuadTree(imageSource string) (*QuadTree, error) {
	img, err := LoadImage(imageSource)
	if err != nil {
		return nil, err
	}

	qt := QuadTree{}

	qt.Width = (*img).Bounds().Max.X - 1 - (*img).Bounds().Min.X
	qt.Height = (*img).Bounds().Max.Y - 1 - (*img).Bounds().Min.Y

	qt.Root = buildQuadTree(img, (*img).Bounds().Min.X, (*img).Bounds().Min.Y, qt.Width, qt.Height)

	return &qt, nil
}

func buildQuadTree(img *image.Image, x, y, w, h int) *QuadTreeNode {
	if w == 0 {
		w = 1
	}
	if h == 0 {
		h = 1
	}

	if w == 1 && h == 1 {
		qn := NewQuadTreeNode((*img).At(x, y))
		return &qn
	}

	qn := QuadTreeNode{}

	qn.NW = buildQuadTree(img, x, y, w/2, h/2)
	qn.NE = buildQuadTree(img, x+(w/2), y, w/2, h/2)
	qn.SW = buildQuadTree(img, x, y+(h/2), w/2, h/2)
	qn.SE = buildQuadTree(img, x+(w/2), y+(h/2), w/2, h/2)

	if qn.NE != nil && qn.NW != nil && qn.SW != nil && qn.SE != nil {
		var red uint32
		var green uint32
		var blue uint32
		var alpha uint32

		red_, green_, blue_, alpha_ := qn.NW.Color.RGBA()

		red += red_
		green += green_
		blue += blue_
		alpha += alpha_

		red_, green_, blue_, alpha_ = qn.NE.Color.RGBA()

		red += red_
		green += green_
		blue += blue_
		alpha += alpha_

		red_, green_, blue_, alpha_ = qn.SE.Color.RGBA()

		red += red_
		green += green_
		blue += blue_
		alpha += alpha_

		red_, green_, blue_, alpha_ = qn.SW.Color.RGBA()

		red += red_
		green += green_
		blue += blue_
		alpha += alpha_

		red /= 4
		green /= 4
		blue /= 4
		alpha /= 4

		qn.Color = colorlib.RGBA64{R: uint16(red), G: uint16(green), B: uint16(blue), A: uint16(alpha)}
	}

	return &qn
}

func (q *QuadTree) Depth() int {
	return q.Root.depth(0)
}

func (q *QuadTreeNode) depth(initialDepth int) int {
	initialDepth++
	if q.NE != nil {
		initialDepth += q.NE.depth(initialDepth)
	}
	if q.NW != nil {
		initialDepth += q.NW.depth(initialDepth)
	}
	if q.SE != nil {
		initialDepth += q.SE.depth(initialDepth)
	}
	if q.SW != nil {
		initialDepth += q.SE.depth(initialDepth)
	}
	return initialDepth
}

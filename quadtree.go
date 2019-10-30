package main

import (
	"image"
	colorlib "image/color"
)

const (
	NE = iota
	NW = iota
	SW = iota
	SE = iota
)

type QuadTreeNode struct {
	Quadrant [4]*QuadTreeNode
	Color    uint32
}

type QuadTree struct {
	Root   *QuadTreeNode
	Height int
	Width  int
}

func BuildQuadTree(img image.Image) (*QuadTree, error) {
	qt := QuadTree{}

	qt.Width = img.Bounds().Max.X - 1 - img.Bounds().Min.X
	qt.Height = img.Bounds().Max.Y - 1 - img.Bounds().Min.Y

	qt.Root = buildQuadTree(img, img.Bounds().Min.X, img.Bounds().Min.Y, qt.Width, qt.Height)

	return &qt, nil
}

func buildQuadTree(img image.Image, x, y, w, h int) *QuadTreeNode {
	if w == 0 && h == 0 {
		qn := QuadTreeNode{Color: PackColor(img.At(x, y))}
		return &qn
	}

	qn := QuadTreeNode{}

	qn.Quadrant[NW] = buildQuadTree(img, x, y, w/2, h/2)
	qn.Quadrant[NE] = buildQuadTree(img, x+(w/2), y, w/2, h/2)
	qn.Quadrant[SW] = buildQuadTree(img, x, y+(h/2), w/2, h/2)
	qn.Quadrant[SE] = buildQuadTree(img, x+(w/2), y+(h/2), w/2, h/2)

	if qn.Quadrant[NE] != nil && qn.Quadrant[NW] != nil && qn.Quadrant[SW] != nil && qn.Quadrant[SE] != nil {
		var red uint32
		var green uint32
		var blue uint32
		var alpha uint32

		for _, child := range []*QuadTreeNode{qn.Quadrant[NE], qn.Quadrant[NW], qn.Quadrant[SE], qn.Quadrant[SW]} {
			red_, green_, blue_, alpha_ := UnpackColor(child.Color).RGBA()
			red += red_
			green += green_
			blue += blue_
			alpha += alpha_
		}

		red /= 4
		green /= 4
		blue /= 4
		alpha /= 4

		qn.Color = PackColor(colorlib.RGBA64{R: uint16(red), G: uint16(green), B: uint16(blue), A: uint16(alpha)})
	}

	return &qn
}

func (q *QuadTree) Leaves() int {
	return q.Root.leaves()
}

func (q *QuadTreeNode) leaves() int {
	leaves := 1
	if q.Quadrant[NE] != nil { //not a leaf node
		leaves += q.Quadrant[NE].leaves()
		leaves += q.Quadrant[NW].leaves()
		leaves += q.Quadrant[SE].leaves()
		leaves += q.Quadrant[SW].leaves()
	}
	return leaves
}

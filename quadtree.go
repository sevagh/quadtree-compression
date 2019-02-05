package main

import (
	"image"
	colorlib "image/color"
)

type QuadTreeNode struct {
	NW    *QuadTreeNode
	NE    *QuadTreeNode
	SE    *QuadTreeNode
	SW    *QuadTreeNode
	Color uint64
}

type QuadTree struct {
	Root   *QuadTreeNode
	Height int
	Width  int
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
	qt.Root.prune()

	return &qt, nil
}

func buildQuadTree(img *image.Image, x, y, w, h int) *QuadTreeNode {
	if w == 0 && h == 0 {
		qn := QuadTreeNode{Color: InterleaveColor((*img).At(x, y))}
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

		for _, child := range []*QuadTreeNode{qn.NE, qn.NW, qn.SE, qn.SW} {
			red_, green_, blue_, alpha_ := DeinterleaveZOrderRGB(child.Color).RGBA()
			red += red_
			green += green_
			blue += blue_
			alpha += alpha_
		}

		red /= 4
		green /= 4
		blue /= 4
		alpha /= 4

		qn.Color = InterleaveColor(colorlib.RGBA64{R: uint16(red), G: uint16(green), B: uint16(blue), A: uint16(alpha)})
	}

	return &qn
}

func (q *QuadTreeNode) prune() {
	if q.NE == nil && q.NW == nil && q.SE == nil && q.SW == nil {
		return
	}

	if q.NE.canPrune(q) && q.NW.canPrune(q) && q.SE.canPrune(q) && q.SW.canPrune(q) {
		q.NE = nil
		q.NW = nil
		q.SE = nil
		q.SW = nil
	} else {
		q.NE.prune()
		q.NW.prune()
		q.SE.prune()
		q.SW.prune()
	}
}

// stack recursion - essentially a DFS of sorts
func (q *QuadTreeNode) canPrune(parent *QuadTreeNode) bool {
	if q.NE == nil { // leaf node
		return CIE76(parent.Color, q.Color) <= 2.3
	}

	return q.NE.canPrune(parent) && q.NW.canPrune(parent) && q.SE.canPrune(parent) && q.SW.canPrune(parent)
}

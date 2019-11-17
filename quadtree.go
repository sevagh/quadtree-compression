package main

import (
	"image"
	colorlib "image/color"
	"github.com/sevagh/k-ary-tree"
)

const (
	NE = iota
	NW = iota
	SW = iota
	SE = iota
)

type QuadTree struct {
	Root   *karytree.Node
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

func buildQuadTree(img image.Image, x, y, w, h int) *karytree.Node {
	qn := karytree.NewNode(nil)

	if w == 0 && h == 0 {
		qn.SetKey(PackColor(img.At(x, y)))
		return &qn
	}

	qn.SetNthChild(NW, buildQuadTree(img, x, y, w/2, h/2))
	qn.SetNthChild(NE, buildQuadTree(img, x+(w/2), y, w/2, h/2))
	qn.SetNthChild(SW, buildQuadTree(img, x, y+(h/2), w/2, h/2))
	qn.SetNthChild(SE, buildQuadTree(img, x+(w/2), y+(h/2), w/2, h/2))

	if qn.NthChild(NE) != nil && qn.NthChild(NW) != nil && qn.NthChild(SW) != nil && qn.NthChild(SE) != nil {
		var red uint32
		var green uint32
		var blue uint32
		var alpha uint32

		for child := uint16(0); child < uint16(4); child++ {
			red_, green_, blue_, alpha_ := UnpackColor(qn.NthChild(child).Key().(uint32)).RGBA()
			red += red_
			green += green_
			blue += blue_
			alpha += alpha_
		}

		red /= 4
		green /= 4
		blue /= 4
		alpha /= 4

		qn.SetKey(PackColor(colorlib.RGBA64{R: uint16(red), G: uint16(green), B: uint16(blue), A: uint16(alpha)}))
	}

	return &qn
}

func (q *QuadTree) Leaves() int {
	return leaves(q.Root)
}

func leaves(q *karytree.Node) int {
	leavesCount := 1
	if q.NthChild(NE) != nil { //not a leaf node
		leavesCount += leaves(q.NthChild(NE))
		leavesCount += leaves(q.NthChild(NW))
		leavesCount += leaves(q.NthChild(SE))
		leavesCount += leaves(q.NthChild(SW))
	}
	return leavesCount
}

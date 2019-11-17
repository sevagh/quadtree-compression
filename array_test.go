package main

import (
	"image/color"
	"reflect"
	"testing"
	"github.com/sevagh/k-ary-tree"
)

func TestQuadTreeObjectToArrayAndBack(t *testing.T) {
	qt := QuadTree{}

	qnNE := karytree.NewNode(PackColor(color.RGBA{R: 255, G: 0, B: 0, A: 0}))
	qnNW := karytree.NewNode(PackColor(color.RGBA{R: 255, G: 0, B: 0, A: 0}))
	qnSE := karytree.NewNode(PackColor(color.RGBA{R: 255, G: 0, B: 0, A: 0}))
	qnSW := karytree.NewNode(PackColor(color.RGBA{R: 255, G: 0, B: 0, A: 0}))

	qn := karytree.NewNode(PackColor(color.RGBA{R: 255, G: 0, B: 0, A: 0}))
	qn.SetNthChild(NE, &qnNE)
	qn.SetNthChild(NW, &qnNW)
	qn.SetNthChild(SE, &qnSE)
	qn.SetNthChild(SW, &qnSW)

	qt.Root = &qn
	qt.Height = 4
	qt.Width = 4

	qtSlice := qt.TreeToSlice()
	qt2 := SliceToTree(&qtSlice)
	qt2Slice := qt2.TreeToSlice()

	if !reflect.DeepEqual(qtSlice, qt2Slice) {
		t.Errorf("Same tree serializes to different array representations: %+v - %+v", qtSlice, qt2Slice)
	}
}

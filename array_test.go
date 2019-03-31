package main

import (
	"image/color"
	"reflect"
	"testing"
)

func TestQuadTreeObjectToArrayAndBack(t *testing.T) {
	qt := QuadTree{}

	qnNE := QuadTreeNode{}
	qnNE.Color = PackColor(color.RGBA{R: 255, G: 0, B: 0, A: 0})

	qnNW := QuadTreeNode{}
	qnNW.Color = PackColor(color.RGBA{R: 255, G: 0, B: 0, A: 0})

	qnSE := QuadTreeNode{}
	qnSE.Color = PackColor(color.RGBA{R: 255, G: 0, B: 0, A: 0})

	qnSW := QuadTreeNode{}
	qnSW.Color = PackColor(color.RGBA{R: 255, G: 0, B: 0, A: 0})

	qn := QuadTreeNode{}
	qn.NE = &qnNE
	qn.NW = &qnNW
	qn.SE = &qnSE
	qn.SW = &qnSW

	qn.Color = PackColor(color.RGBA{R: 255, G: 0, B: 0, A: 0})

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

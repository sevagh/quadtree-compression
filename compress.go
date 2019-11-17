package main

import (
	"image"
	"github.com/sevagh/k-ary-tree"
)

func (q *QuadTree) Compress(level int) (*QuadTree, error) {
	img := image.NewNRGBA(image.Rect(0, 0, q.Width, q.Height))

	for y := 0; y < q.Height; y++ {
		for x := 0; x < q.Width; x++ {
			img.Set(x, y, UnpackColor(q.getPixel(x, y, level)))
		}
	}

	return BuildQuadTree(img)
}

func (q *QuadTree) getPixel(x, y, level int) uint32 {
	if q.Root != nil {
		return getPixel(q.Root, x, y, q.Width, q.Height, 0, level)
	}
	return 0
}

func getPixel(qn *karytree.Node, x, y, xCoord, yCoord, level, maxLevel int) uint32 {
	level += 1
	if maxLevel > 0 && level >= maxLevel {
		return qn.Key().(uint32)
	}
	if qn.NthChild(NW) != nil && x < xCoord/2 && y < yCoord/2 {
		return getPixel(qn.NthChild(NW), x, y, xCoord/2, yCoord/2, level, maxLevel)
	}
	if qn.NthChild(NE) != nil && x >= xCoord/2 && y < yCoord/2 {
		return getPixel(qn.NthChild(NE), x-xCoord/2, y, xCoord/2, yCoord/2, level, maxLevel)
	}
	if qn.NthChild(SW) != nil && x < xCoord/2 && y >= yCoord/2 {
		return getPixel(qn.NthChild(SW), x, y-yCoord/2, xCoord/2, yCoord/2, level, maxLevel)
	}
	if qn.NthChild(SE) != nil && x >= xCoord/2 && y >= yCoord/2 {
		return getPixel(qn.NthChild(SE), x-xCoord/2, y-yCoord/2, xCoord/2, yCoord/2, level, maxLevel)
	}
	return qn.Key().(uint32)
}

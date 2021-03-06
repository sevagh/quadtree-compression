package main

import "image"

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
		return q.Root.getPixel(x, y, q.Width, q.Height, 0, level)
	}
	return 0
}

func (qn *QuadTreeNode) getPixel(x, y, xCoord, yCoord, level, maxLevel int) uint32 {
	level += 1
	if maxLevel > 0 && level >= maxLevel {
		return qn.Color
	}
	if qn.Quadrant[NW] != nil && x < xCoord/2 && y < yCoord/2 {
		return qn.Quadrant[NW].getPixel(x, y, xCoord/2, yCoord/2, level, maxLevel)
	}
	if qn.Quadrant[NE] != nil && x >= xCoord/2 && y < yCoord/2 {
		return qn.Quadrant[NE].getPixel(x-xCoord/2, y, xCoord/2, yCoord/2, level, maxLevel)
	}
	if qn.Quadrant[SW] != nil && x < xCoord/2 && y >= yCoord/2 {
		return qn.Quadrant[SW].getPixel(x, y-yCoord/2, xCoord/2, yCoord/2, level, maxLevel)
	}
	if qn.Quadrant[SE] != nil && x >= xCoord/2 && y >= yCoord/2 {
		return qn.Quadrant[SE].getPixel(x-xCoord/2, y-yCoord/2, xCoord/2, yCoord/2, level, maxLevel)
	}
	return qn.Color
}

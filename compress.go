package main

import (
	"fmt"
	"image/color"
)

func (q *QuadTree) Compress(tolerance float64) {
	q.root.compress(tolerance)
}

func (q *quadTreeNode) compress(tolerance float64) {
	if q.children[NE] != nil && q.children[NW] != nil && q.children[SE] != nil && q.children[SW] != nil {
		thisColor := q.color

		childrenColor := make([]color.Color, 4)

		childrenColor[NE] = q.children[NE].color
		childrenColor[NW] = q.children[NW].color
		childrenColor[SW] = q.children[SW].color
		childrenColor[SE] = q.children[SE].color

		allBelowJND := true

		for i := 0; i < 4; i++ {
			colorDist := CIE76(thisColor, childrenColor[i])
			if colorDist > 500 {
				allBelowJND = false
				break
			} else {
				fmt.Printf("GOT A LOW ONE %f\n", colorDist)
			}
		}

		if allBelowJND {
			fmt.Printf("CLEARED SOME!\n")
			// the 4 children are within Just Noticeable Difference of parent
			// and can be compressed
			for i := 0; i < 4; i++ {
				q.children[i] = nil
			}
		} else {
			q.children[NE].compress(tolerance)
			q.children[NW].compress(tolerance)
			q.children[SW].compress(tolerance)
			q.children[SE].compress(tolerance)
		}
	}
}

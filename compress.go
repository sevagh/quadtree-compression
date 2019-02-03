package main

func (q *QuadTree) Compress(tolerance float64) {
	q.root.prune(tolerance)
}

func (q *quadTreeNode) prune(tolerance float64) {
	if q.children[NE] == nil && q.children[NW] == nil && q.children[SE] == nil && q.children[SW] == nil {
		return
	}

	if q.children[NE].canPrune(q, tolerance) && q.children[NW].canPrune(q, tolerance) && q.children[SE].canPrune(q, tolerance) && q.children[SW].canPrune(q, tolerance) {
		q.children[NE] = nil
		q.children[NW] = nil
		q.children[SE] = nil
		q.children[SW] = nil
	} else {
		q.children[NE].prune(tolerance)
		q.children[NW].prune(tolerance)
		q.children[SE].prune(tolerance)
		q.children[SW].prune(tolerance)
	}
}

// stack recursion - essentially a DFS of sorts
func (q *quadTreeNode) canPrune(parent *quadTreeNode, tolerance float64) bool {
	if q.children[NE] == nil || q.children[NW] == nil || q.children[SE] == nil || q.children[SW] == nil { // leaf node
		//colorDist := CIE76(parent.color, q.color)
		colorDist := EuclidianDistance(parent.color, q.color)

		// can prune if the LAB CIE76 is under the Just Noticeable Difference threshold
		return colorDist <= tolerance
	}

	return q.children[NE].canPrune(parent, tolerance) && q.children[NW].canPrune(parent, tolerance) && q.children[SE].canPrune(parent, tolerance) && q.children[SW].canPrune(parent, tolerance)
}

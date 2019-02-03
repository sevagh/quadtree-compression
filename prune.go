package main

func (q *QuadTree) Prune() {
	q.Root.prune()
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
	if q.NE == nil || q.NW == nil || q.SE == nil || q.SW == nil { // leaf node
		colorDist := CIE76(parent.Color, q.Color)

		// can prune if the LAB CIE76 is under the Just Noticeable Difference threshold
		return colorDist <= 2.3
	}

	return q.NE.canPrune(parent) && q.NW.canPrune(parent) && q.SE.canPrune(parent) && q.SW.canPrune(parent)
}

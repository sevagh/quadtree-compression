package main

import "github.com/sevagh/quadtree-compression/k-ary-tree"

func (q *QuadTree) TreeToSlice() []uint32 {
	treeArrRep := make([]uint32, q.Leaves()+2)

	treeArrRep[0] = uint32(q.Root.Key()) //root = index 0

	DFS(q.Root, 0, &treeArrRep)

	treeArrRep[len(treeArrRep)-2] = uint32(q.Width)
	treeArrRep[len(treeArrRep)-1] = uint32(q.Height)

	return treeArrRep
}

func SliceToTree(nodes *[]uint32) QuadTree {
	qt := QuadTree{}

	qt.Height = int((*nodes)[len(*nodes)-1])
	qt.Width = int((*nodes)[len(*nodes)-2])

	qn := karytree.NewNode(0)
	qn.SetKey((*nodes)[0])

	qt.Root = &qn

	InvDFS(qt.Root, 0, nodes)
	return qt
}

func InvDFS(q *karytree.Node, index int, storage *[]uint32) {
	// no more nodes - use 4*index+3 to omit height/width at the end
	if 4*index+3 >= len(*storage) {
		return
	}

	qNE := karytree.NewNode((*storage)[4*index+1])
	q.SetNthChild(NE, &qNE)
	InvDFS(&qNE, 4*index+1, storage)

	qNW := karytree.NewNode((*storage)[4*index+2])
	q.SetNthChild(NW, &qNW)
	InvDFS(&qNW, 4*index+2, storage)

	qSE := karytree.NewNode((*storage)[4*index+3])
	q.SetNthChild(SE, &qSE)
	InvDFS(&qSE, 4*index+3, storage)

	qSW := karytree.NewNode((*storage)[4*index+4])
	q.SetNthChild(SW, &qSW)
	InvDFS(&qSW, 4*index+4, storage)
}

//store the tree as follows
// NE: 4*i + 1
// NW: 4*i + 2
// SE: 4*i + 3
// SW: 4*i + 4
func DFS(q *karytree.Node, index int, storage *[]uint32) {
	if q.NthChild(NE) != nil { //not a leaf node
		(*storage)[4*index+1] = uint32(q.NthChild(NE).Key())
		DFS(q.NthChild(NE), 4*index+1, storage)

		(*storage)[4*index+2] = uint32(q.NthChild(NW).Key())
		DFS(q.NthChild(NW), 4*index+2, storage)

		(*storage)[4*index+3] = uint32(q.NthChild(SE).Key())
		DFS(q.NthChild(SE), 4*index+3, storage)

		(*storage)[4*index+4] = uint32(q.NthChild(SW).Key())
		DFS(q.NthChild(SW), 4*index+4, storage)
	}
}

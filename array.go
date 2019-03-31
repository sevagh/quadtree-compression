package main

func (q *QuadTree) TreeToSlice() []uint32 {
	treeArrRep := make([]uint32, q.Leaves()+2)

	treeArrRep[0] = uint32(q.Root.Color) //root = index 0

	DFS(q.Root, 0, &treeArrRep)

	treeArrRep[len(treeArrRep)-2] = uint32(q.Width)
	treeArrRep[len(treeArrRep)-1] = uint32(q.Height)

	return treeArrRep
}

func SliceToTree(nodes *[]uint32) QuadTree {
	qt := QuadTree{}

	qt.Height = int((*nodes)[len(*nodes)-1])
	qt.Width = int((*nodes)[len(*nodes)-2])

	qn := QuadTreeNode{}
	qn.Color = (*nodes)[0]

	qt.Root = &qn

	InvDFS(qt.Root, 0, nodes)
	return qt
}

func InvDFS(q *QuadTreeNode, index int, storage *[]uint32) {
	// no more nodes - use 4*index+3 to omit height/width at the end
	if 4*index+3 >= len(*storage) {
		return
	}

	qNE := QuadTreeNode{Color: (*storage)[4*index+1]}
	q.NE = &qNE
	InvDFS(q.NE, 4*index+1, storage)

	qNW := QuadTreeNode{Color: (*storage)[4*index+2]}
	q.NW = &qNW
	InvDFS(q.NW, 4*index+2, storage)

	qSE := QuadTreeNode{Color: (*storage)[4*index+3]}
	q.SE = &qSE
	InvDFS(q.SE, 4*index+3, storage)

	qSW := QuadTreeNode{Color: (*storage)[4*index+4]}
	q.SW = &qSW
	InvDFS(q.SW, 4*index+4, storage)
}

//store the tree as follows
// NE: 4*i + 1
// NW: 4*i + 2
// SE: 4*i + 3
// SW: 4*i + 4
func DFS(q *QuadTreeNode, index int, storage *[]uint32) {
	if q.NE != nil { //not a leaf node
		(*storage)[4*index+1] = uint32(q.NE.Color)
		DFS(q.NE, 4*index+1, storage)

		(*storage)[4*index+2] = uint32(q.NW.Color)
		DFS(q.NW, 4*index+2, storage)

		(*storage)[4*index+3] = uint32(q.SE.Color)
		DFS(q.SE, 4*index+3, storage)

		(*storage)[4*index+4] = uint32(q.SW.Color)
		DFS(q.SW, 4*index+4, storage)
	}
}

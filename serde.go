package main

import (
	"encoding/gob"
	"os"
)

func (q *QuadTree) SerializeToFile(path string) error {
	q.Prune()

	qFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer qFile.Close()

	encoder := gob.NewEncoder(qFile)

	treeSlice := q.TreeToSlice()
	return encoder.Encode(treeSlice)
}

func LoadQuadTreeFromFile(path string) (QuadTree, error) {
	var treeSlice []int64
	qFile, err := os.Open(path)
	if err != nil {
	}
	defer qFile.Close()

	decoder := gob.NewDecoder(qFile)
	err = decoder.Decode(&treeSlice)

	q := SliceToTree(&treeSlice)
	return q, err
}

func (q *QuadTree) TreeToSlice() []int64 {
	treeArrRep := []int64{}

	treeArrRep = append(treeArrRep, int64(q.Root.Color)) //root = index 0

	initialIndex := 0
	DFS(q.Root, initialIndex, &treeArrRep)

	treeArrRep = append(treeArrRep, int64(q.Width))
	treeArrRep = append(treeArrRep, int64(q.Height))

	return treeArrRep
}

func SliceToTree(nodes *[]int64) QuadTree {
	qt := QuadTree{}

	qt.Height = int((*nodes)[len(*nodes)-1])
	qt.Width = int((*nodes)[len(*nodes)-2])

	*nodes = (*nodes)[:len(*nodes)-2]

	qn := QuadTreeNode{}
	qn.Color = PackedRGB((*nodes)[0])

	qt.Root = &qn

	InvDFS(qt.Root, 0, nodes)
	return qt
}

func InvDFS(q *QuadTreeNode, index int, storage *[]int64) {
	// still have more nodes
	if 2*index+1 < len(*storage) {
		qNE := NewQuadTreeNode(PackedRGB((*storage)[2*index+1]))
		q.NE = &qNE
		InvDFS(q.NE, 2*index+1, storage)

		qNW := NewQuadTreeNode(PackedRGB((*storage)[2*index+2]))
		q.NW = &qNW
		InvDFS(q.NW, 2*index+2, storage)

		qSE := NewQuadTreeNode(PackedRGB((*storage)[2*index+3]))
		q.SE = &qSE
		InvDFS(q.SE, 2*index+3, storage)

		qSW := NewQuadTreeNode(PackedRGB((*storage)[2*index+4]))
		q.SW = &qSW
		InvDFS(q.SW, 2*index+4, storage)
	}
}

//store the tree as follows
// NE: 2*i + 1
// NW: 2*i + 2
// SE: 2*i + 3
// SW: 2*i + 4
func DFS(q *QuadTreeNode, index int, storage *[]int64) {
	if q.NE != nil { //not a leaf node
		*storage = append(*storage, int64(q.NE.Color))
		DFS(q.NE, 2*index+1, storage)
		*storage = append(*storage, int64(q.NW.Color))
		DFS(q.NW, 2*index+2, storage)
		*storage = append(*storage, int64(q.NE.Color))
		DFS(q.SE, 2*index+3, storage)
		*storage = append(*storage, int64(q.SW.Color))
		DFS(q.SW, 2*index+4, storage)
	}
}

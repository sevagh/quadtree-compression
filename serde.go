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

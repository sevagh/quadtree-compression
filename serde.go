package main

import (
	"encoding/gob"
	"image/color"
	"os"
)

func init() {
	gob.Register(color.RGBA64{})
}

func (q *QuadTree) SerializeToFile(path string) error {
	qFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer qFile.Close()

	encoder := gob.NewEncoder(qFile)
	return encoder.Encode(q)
}

func LoadQuadTreeFromFile(path string) (QuadTree, error) {
	var q QuadTree
	qFile, err := os.Open(path)
	if err != nil {
	}
	defer qFile.Close()

	decoder := gob.NewDecoder(qFile)
	err = decoder.Decode(&q)
	return q, err
}

package quadtree_compression

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"io"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/sevagh/quadtree-compression/quadtree_proto"
)

// chose zlib + protobuf from https://eng.uber.com/trip-data-squeeze/

func (q *QuadTree) SerializeToFile(path string) error {
	q.Prune()

	treeSlice := q.TreeToSlice()
	protoObj := quadtree_proto.ImageQuadtree{Data: treeSlice}

	data, err := proto.Marshal(&protoObj)
	if err != nil {
		return err
	}

	var compressedData bytes.Buffer
	w := zlib.NewWriter(&compressedData)
	defer w.Close()

	w.Write(data)
	w.Flush()

	writeFile, err := os.Create(path)
	if err != nil {
		return err
	}

	_, err = io.Copy(writeFile, bufio.NewReader(&compressedData))
	if err != nil {
		return err
	}

	return nil
}

func LoadQuadTreeFromFile(path string) (QuadTree, error) {
	readFile, err := os.Open(path)
	if err != nil {
		return QuadTree{}, err
	}

	r, err := zlib.NewReader(readFile)
	defer r.Close()

	var protoObj quadtree_proto.ImageQuadtree

	uncompressedData := new(bytes.Buffer)
	uncompressedData.ReadFrom(r)

	err = proto.Unmarshal(uncompressedData.Bytes(), &protoObj)
	if err != nil {
		return QuadTree{}, err
	}

	q := SliceToTree(&protoObj.Data)
	return q, nil
}

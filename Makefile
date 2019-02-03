proto:
	protoc --go_out=quadtree_proto --proto_path=quadtree_proto QuadTree.proto

test: proto
	go test -v ./...

fmt:
	go fmt ./...

build:
	go build ./

clean:
	@rm -rf *.png *.quadtree quadtree-compression

.PHONY: proto test fmt clean build

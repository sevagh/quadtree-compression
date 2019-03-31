proto:
	protoc --go_out=proto --proto_path=proto QuadTree.proto

all: build

test: proto
	go test -v ./...

fmt:
	go fmt ./...

build: proto
	go build ./

clean:
	@rm -rf *.png *.quadtree *.gif

.PHONY: proto test fmt clean build

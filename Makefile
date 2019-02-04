proto:
	protoc --go_out=quadtree_proto --proto_path=quadtree_proto QuadTree.proto

test: proto
	go test -v ./...

fmt:
	go fmt ./...

build:
	go build ./

gif_demo:
	@mkdir -p bin/
	go build -o bin/gif_demo ./gif_demo/

clean:
	@rm -rf *.png *.quadtree *.gif

.PHONY: proto test fmt clean build gif_demo

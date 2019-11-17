all: proto

proto:
	protoc --go_out=proto --proto_path=proto QuadTree.proto

fmt:
	go fmt ./...

test:
	go test -v ./...

bench:
	go test -run=^a -bench=Bench -v -benchmem


build:
	go build .

clean:
	@rm -rf *.png *.quadtree *.gif

.PHONY: proto clean

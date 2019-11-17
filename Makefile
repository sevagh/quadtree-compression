all: proto

proto:
	protoc --go_out=proto --proto_path=proto QuadTree.proto

bench:
	go test -run=^a -bench=Bench -v -benchmem

clean:
	@rm -rf *.png *.quadtree *.gif

.PHONY: proto clean

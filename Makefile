all: proto

proto:
	protoc --go_out=proto --proto_path=proto QuadTree.proto

clean:
	@rm -rf *.png *.quadtree *.gif

.PHONY: proto clean

all: build

test:
	go test -v ./...
	@make clean

fmt:
	go fmt ./...

build:
	go build ./

clean:
	@rm -rf *.png *.quadtree *.gif

.PHONY: proto test fmt clean build

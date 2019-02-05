all: build

test:
	go test -v ./...

fmt:
	go fmt ./...

build:
	go build ./

clean:
	@rm -rf *.png *.quadtree *.gif

.PHONY: proto test fmt clean build

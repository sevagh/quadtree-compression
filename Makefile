all: build

GO_TEST_FLAGS:=-race -count=1 -v
GO_BENCH_FLAGS:=-benchmem -v -cpuprofile=cpuprofile.out -memprofile=memprofile.out -run=^a
run?=""
runarg?=-run=
component?=./

coverage: GO_TEST_FLAGS:=$(GO_TEST_FLAGS) -coverprofile=coverage.out -cover
coverage: test
	go tool cover -html=coverage.out

bench: GO_TEST_FLAGS:=$(GO_BENCH_FLAGS)
bench: run:=Bench
bench: runarg:=-bench=
bench: test

test:
	go test $(GO_TEST_FLAGS) $(component) $(runarg)$(run)

proto:
	protoc --go_out=proto --proto_path=proto QuadTree.proto

build: proto
build:
	go build -mod=vendor .

clean:
	@rm -rf *.png *.quadtree *.gif

fmt:
	go fmt $(component)

lint:
	-go vet $(component)

.PHONY: proto clean

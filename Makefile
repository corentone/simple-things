all: bench
build:
	go build ./...
test:
	go test ./...
bench:
	go test -bench=. ./...
deps:
	go get -u github.com/stretchr/testify
.PHONY: test

build:
	go build -o main .

bench: build
	go test -bench . -benchtime 100x
	go test -bench . -benchtime 1000x
	go test -bench . -benchtime 10000x

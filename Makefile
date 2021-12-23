.DEFAULT_GOAL := build

build:
	go build -o bin/display-placer-helper cmd/main.go

run:
	go run cmd/main.go
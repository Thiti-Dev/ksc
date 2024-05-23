.PHONY: build-window build-osx build-linux

build-window:
	GOOS=windows go build -o ksc ./cmd/cli/*
build-osx:
	GOOS=darwin go build -o ksc ./cmd/cli/*
build-linux:
	GOOS=linux go build -o ksc ./cmd/cli/*
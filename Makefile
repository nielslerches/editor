.PHONY: build run

SRC := $(wildcard **/*.go)

all: Makefile build

build: main

main: main.go $(SRC)
	go build -o main main.go

run:
	go run main.go

clean:
	rm main

.PHONY: default build exec

FILES := $(wildcard *.go)
BIN   := filebrowse.exe

default: build run

build: $(FILES)
	go build -ldflags="-H windowsgui" -o $(BIN)

run:
	./$(BIN)

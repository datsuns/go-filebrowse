.PHONY: default build exec

FILES := $(wildcard *.go)
BIN   := filebrowse.exe

default: build run

build: $(FILES)
	@echo go build -ldflags="-H windowsgui" -o $(BIN)
	go build  -o $(BIN)

run:
	./$(BIN)

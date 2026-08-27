.PHONY: build run clean

build:
	@go build -o bin/api ./cmd/api

run: build
	@./bin/api

clean:
	go clean
	-del /Q bin\main.exe bin\main 2>nul

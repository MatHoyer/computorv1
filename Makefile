all : build

build:
	@go build -o bin/main ./src

run:
	@go run ./src
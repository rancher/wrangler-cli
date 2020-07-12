all: validate build build-example

build-example:
	go build -o bin/example ./example

build:
	go build ./...

validate:
	go fmt
	go vet

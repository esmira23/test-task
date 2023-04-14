run:
	go run main.go

build:
	go build

test: 
	go test -v

.PHONY: run build test 
.PHONY: deps dev test build

deps:
	go get -u ./...

build: deps
	docker-compose up --build

dev:
	docker-compose up

test:
	go test -v ./...

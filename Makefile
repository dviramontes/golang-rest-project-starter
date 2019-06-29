.PHONY: deps dev test

deps:
	go get -u ./...

dev: deps
	docker-compose up --build

test:
	go test -v ./...

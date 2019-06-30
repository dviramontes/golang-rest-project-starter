.PHONY: deps dev test build tml

deps:
	go get -u ./...

build: deps
	docker-compose up --build

dev:
	docker-compose up

test:
	go test -v ./...

tml:
	cd tml
	qtc

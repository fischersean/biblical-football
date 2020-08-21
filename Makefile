.PHONY: run image install build clean vueui

vueui:
	vue ui -p 8081

clean:
	rm app
	rm popdb

build:
	go build ./cmd/app

run:
	./app

up:
	make build && make run

image:
	docker-compose up --build

install:
	go install -v ./cmd/app

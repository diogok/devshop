all: build

build: deps
	docker build -t diogok/devshop .

push:
	docker push diogok/devshop

run:
	docker-compose up

deps:
	docker-compose run app sh -c "cat /app/src/*.go | grep 'github.com\\/[^\"]*' -o | xargs go get"

tests:
	docker-compose run app go test src/*.go

test: tests


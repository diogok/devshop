all: build

build: deps
	docker build -t diogok/devshop .

push:
	docker push diogok/devshop

run-dev:
	docker-compose run -p 8080:8080 app /app/run.sh dev

run:
	docker-compose up

deps:
	docker-compose run app sh -c "cat /app/src/*.go | grep 'github.com\\/[^\"]*' -o | xargs go get"
	docker-compose run app sh -c "cat /app/src/*.go | grep 'golang.org\\/[^\"]*' -o | xargs go get"

tests:
	docker-compose run app go test -v src/*.go

test: tests


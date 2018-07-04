SERVICE=backend	

build-deps:
	$(MAKE) -C ./$(SERVICE) MAKEFLAGS=build-deps

push:
	docker-compose push

build:
	docker-compose build

build-with-deps: build-deps
	docker-compose build --no-cache

run: up

start: up

up:
	docker-compose up -d

stop: down

down:
	docker-compose down

restart:
	docker-compose restart

rm:
	docker-compose rm -f

log: logs

logs:
	docker-compose logs

envs:
	docker-compose run $(SERVICE) env

enter:
	docker-compose run $(SERVICE) /bin/sh

test:
	docker-compose run $(SERVICE) /bin/sh -c "go test ./src/..."

test-local:
	$(MAKE) -C ./$(SERVICE) MAKEFLAGS=test
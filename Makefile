
### Edit these variables ###
IMG_NAME=go-docker
PORT=80
TAG=latest
### End of edit ###

IMG=whimatthew/$(IMG_NAME)
CONTAINER=$(IMG_NAME)
RUNOPTS=-p $(PORT):80
FQ_IMG?=$(IMG):$(TAG)

SERVICE=web
TEST_REGEX=*test.py

push:
	docker-compose push

build: build-go
	docker-compose build

build-go: set-envs
	go build -a -installsuffix cgo -o ./web/go/bin/app.a ./web/go/src/server

set-envs:
ifeq ($(findstring Windows,$(OS)),Windows)
	set CGO_ENABLED=0
	set GOOS=linux
	set GOARCH=amd64
else
	CGO_ENABLED=0
	GOOS=linux
	GOARCH=amd64
endif

up:
	docker-compose up -d

down:
	docker-compose down

restart:
	docker-compose restart

rm:
	docker-compose rm -f

logs:
	docker-compose logs

envs:
	docker-compose run $(SERVICE) env

enter:
	docker-compose run $(SERVICE) bash

test: build
	docker-compose run $(SERVICE) bash -c "/usr/local/bin/python -m unittest discover --pattern $(TEST_REGEX)"


### Edit these variables ###
IMG_NAME=go-docker
PORT=80
TAG=latest
### End of edit ###

IMG=whimatthew/$(IMG_NAME)
CONTAINER=$(IMG_NAME)
RUNOPTS=-p $(PORT):80
FQ_IMG?=$(IMG):$(TAG)

SERVICE=app
TEST_REGEX=*test.py

push:
	docker-compose push

build:
	docker-compose build

up:
	docker-compose up -d

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
	docker-compose run $(SERVICE) /bin/bash

# TODO test
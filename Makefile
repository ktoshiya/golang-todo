.PHONY: help build build-loacl up down logs ps test
.DEFAULT_GOAL := help

DOCKER_TAG := latest

build:
	docker build -t ktoshiya/gotodo:${DOCKER_TAG} \
			--target deploy ./

build-loacl:
	docker compose build --no-cache

up:
	docker compose up -d

down:
	docker compose down

logs:
	docker compose logs -f

ps:
	docker compose ps

test:
	go test -race -shuffle=on ./...

help:

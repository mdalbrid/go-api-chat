.PHONY: help

PROJECTNAME=$(shell basename "$(PWD)")

## build ...
build:
	go build -v ./cmd/main/app.go

## run ...
run:
	go run ./cmd/main/app.go

## test ...
test:
	go test -v ./...

## migrate_up ...
migrate_up:
	migrate -path ./schema -database 'postgres://postgres:12345@localhost:5432/postgres?sslmode=disable' up

## migrate_down ...
migrate_down:
	migrate -path ./schema -database 'postgres://postgres:12345@localhost:5432/postgres?sslmode=disable' down

## docker_all ...
docker_all:
	docker-compose build go-chat
	docker-compose up go-chat

## docker_build ...
docker_build:
	docker-compose build go-chat

## docker_up ...
docker_up:
	docker-compose up -d go-chat
	docker-compose ps

## docker_down ...
docker_down:
	docker-compose down
	docker-compose ps

## docker_rm
docker_rm:
	#docker volume rm $$(docker volume ls -q)
	sudo rm -rf $(PWD)/.database

help: Makefile
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/    /'
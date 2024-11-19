latest_tag := $(shell git describe --tags 2> /dev/null || git rev-parse --short HEAD)
new_migration: name ?= $(shell uuidgen)
db_%: uri ?= "postgres://postgres:postgres@localhost:5432/mobility-payments?sslmode=disable"
db_upgrade: n ?=
db_dowgrade: n ?= 1

dev:
	air -c .air.toml

run:
	go run cmd/mobility-payments/main.go

build:
	go build -o bin/mobility-payments cmd/mobility-payments/main.go

docker_build: test write_version
	docker build . -t goyurback:$(latest_tag)

test:
	go test -v ./...

new_migration:
	migrate create -ext sql -dir migrations $(name)

db_upgrade:
	migrate -path migrations -database $(uri) up $(n)

db_dowgrade:
	migrate -path migrations -database $(uri) down $(n)

gen_config:
	pkl-gen-go config/Config.pkl

write_version:
	@echo $(latest_tag) > cmd/mobility-payments/.version

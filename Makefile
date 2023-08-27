include .env

run/dev:
	go run ./cmd/server

migrations/create:
	goose -dir=./data/migrations/ postgres ${DSN} create ${name} sql

migrations/up:
	goose -dir=./data/migrations/ postgres ${DSN} up

migrations/down:
	goose -dir=./data/migrations/ postgres ${DSN} down

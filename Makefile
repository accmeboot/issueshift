include .env

run/dev:
	go run ./cmd/server

migrations/create:
	goose -dir=./migrations/ postgres ${DSN} create ${name} sql

migrations/up:
	goose -dir=./migrations/ postgres ${DSN} up

migrations/down:
	goose -dir=./migrations/ postgres ${DSN} down

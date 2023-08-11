.PHONY: migrate-up .run

migrate-up:
	go run cmd/migrate/main.go up

run:
	go run cmd/server/main.go run
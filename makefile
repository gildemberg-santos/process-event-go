build:
	@docker compose build

down:
	@docker compose down

up:
	@docker compose up

up_d:
	@docker compose up -d

dev:
	@go run cmd/main.go

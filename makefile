build:
	@docker compose build --no-cache

build-dev:
	@docker compose -f docker-compose-dev.yml build --no-cache

down:
	@docker compose down

down-dev:
	@docker compose -f docker-compose-dev.yml down

up:
	@docker compose up

up-dev:
	@docker compose -f docker-compose-dev.yml up -d

dev:
	@rm -rf ./db/development.db && go run cmd/main.go

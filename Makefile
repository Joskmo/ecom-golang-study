include .env
export

.PHONY: migrate-up migrate-down migrate-create

migrate-up:
	goose -dir $(GOOSE_MIGRATIONS_DIR) up

migrate-down:
	goose -dir $(GOOSE_MIGRATIONS_DIR) down

migrate-create:
	goose -dir $(GOOSE_MIGRATIONS_DIR) -s create $(name) sql
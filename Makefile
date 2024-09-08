DB_CONNECTION_STRING="root:password@tcp(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"

.PHONY: migrate-up migrate-down migration

migrate:
	@goose -dir ./database/migrations mysql $(DB_CONNECTION_STRING) up

migrate-down:
	@goose -dir ./database/migrations mysql $(DB_CONNECTION_STRING) down

migrate-fresh:
	@goose -dir ./database/migrations mysql $(DB_CONNECTION_STRING) down-to 0

migration:
	@read -p "Enter migration name: " name; \
	goose -dir ./database/migrations create $$name sql
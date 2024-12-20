db-migrate:
	@echo "Migrating database"
	@go run tools/sql_lite_migration.go

run:
	@echo "Running server"
	@go run main.go
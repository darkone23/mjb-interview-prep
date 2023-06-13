set dotenv-load

setup:
	go install github.com/volatiletech/sqlboiler/v4@v4.14.2
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-sqlite3@v4.14.2
server:
	go run cmd/server/main.go
migrate:
	go run cmd/migrations/main.go
codegen:
	sqlboiler -c $SQL_CONFIG sqlite3
setup-and-run: setup migrate codegen server
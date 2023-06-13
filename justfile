set dotenv-load

setup:
	go install github.com/volatiletech/sqlboiler/v4@v4.14.2
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-sqlite3@v4.14.2
lint:
	staticcheck ./...
server:
	go run cmd/server/main.go
migrate:
	go run cmd/migrations/main.go
codegen:
	sqlboiler -c $SQL_CONFIG sqlite3
	find models -name '*test.go' -delete
ci: setup migrate codegen lint
setup-and-serve: setup migrate codegen server
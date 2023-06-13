server:
	go run cmd/server/main.go
migrate:
	go run cmd/migrations/main.go
codegen:
	sqlboiler sqlite3
setup:
	mkdir -p data && touch data/db.sqlite3
	go install github.com/volatiletech/sqlboiler/v4@v4.14.2
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-sqlite3@latest
setup-and-run: setup migrate codegen server
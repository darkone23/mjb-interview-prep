package db

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func RunMigrations() {
	// Database connection string
	log.Println("Migrations: About to run...")

	dbURL := "./data/db.sqlite3" // "postgres://admin:admin@localhost/test_repo?sslmode=disable"
	dbDriver := "sqlite3"
	dbName := "app"

	db, err := sql.Open(dbDriver, dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a new instance of the PostgreSQL driver for migrate
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://internal/migrations", dbName, driver)
	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	fmt.Println("Database migration complete.")
}

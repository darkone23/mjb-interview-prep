package db

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"os"
)

func RunMigrations(conf SqlConf) {
	// Database connection string

	data_dir := os.Getenv("DATA_HOME")
	dbURL := fmt.Sprintf("file:%s/%s", data_dir, conf.Sqlite.DbName)

	log.Printf("Migrations: About to migrate: %s", dbURL)

	dbDriver := "sqlite3"
	dbName := "sqlite3"

	db, err := sql.Open(dbDriver, dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{
		DatabaseName: dbName,
	})
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

	log.Println("Migrations: complete!")
}

package main

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"

	"github.com/gin-gonic/gin"
	"log"
	"mjb-interview-prep/internal/user"
	"net/http"
)

func main() {

	log.Println("Migrations: About to run...")
	runMigrations()
	log.Println("Migrations: complete!")

	svc, err := user.NewService("admin", "admin")
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	h := user.Handler{Svc: *svc}

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	r.POST("/user", h.AddUser)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	// http.HandleFunc("/user", h.AddUser)

	log.Println("HTTP: starting listen at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func runMigrations() {
	// Database connection string

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

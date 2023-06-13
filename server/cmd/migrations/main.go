package main

import (
	"mjb-interview-prep/internal/db"
)

func main() {

	config := db.LoadConfig()
	db.RunMigrations(config)

}

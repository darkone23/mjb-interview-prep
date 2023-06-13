package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"mjb-interview-prep/internal/db"
	"mjb-interview-prep/internal/user"
)

func main() {

	db.RunMigrations()

	svc, err := user.NewService("admin", "admin")
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	h := user.Handler{Svc: *svc}

	r.POST("/user", h.AddUser)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

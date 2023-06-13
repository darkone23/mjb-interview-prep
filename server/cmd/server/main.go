package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"mjb-interview-prep/internal/db"
	"mjb-interview-prep/internal/user"
	"net/http"
	"os"
)

func main() {

	conf := db.LoadConfig()
	db.RunMigrations(conf)

	svc, err := user.NewService(conf, 64)
	svc.Open()
	defer svc.Close()

	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.SetTrustedProxies(nil)

	h := user.NewUserRoutes(svc)
	r.POST("/api/user", h.AddUser)
	r.GET("/api/user", h.ListUser)
	r.GET("/api/user/:user_id", h.GetUser)

	// catch all route for static assets
	static_dir := os.Getenv("CLIENT_DIST_DIR")
	r.NoRoute(gin.WrapH(http.FileServer(http.Dir(static_dir))))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

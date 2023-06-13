package user

import (
	// "encoding/json"
	// "fmt"
	"log"

	"github.com/gin-gonic/gin"
	// "log"
	"net/http"
)

type Handler struct {
	Svc service
}

func (h Handler) AddUser(c *gin.Context) {
	var user User
	err := c.BindJSON(&user)
	if err != nil {
		log.Printf("No good: %s\n", err)
	}
	message, err := h.Svc.AddUser(user)
	if err != nil {
		log.Printf("Error creating user: %s", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, message)
	}
}

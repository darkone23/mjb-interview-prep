package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type UserRoutes struct {
	svc service
}

func NewUserRoutes(svc service) UserRoutes {
	return UserRoutes{
		svc: svc,
	}
}

func (h UserRoutes) AddUser(c *gin.Context) {
	var user UserDto
	err := c.BindJSON(&user)
	if err != nil {
		log.Printf("No good: %s\n", err)
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		message, err := h.svc.AddUser(user)
		if err != nil {
			log.Printf("Error creating user: %s", err)
			c.AbortWithStatus(http.StatusInternalServerError)
		} else {
			c.JSON(http.StatusOK, message)
		}
	}
}

func (h UserRoutes) ListUser(c *gin.Context) {
	users, err := h.svc.FindAllUsers()
	if err != nil {
		log.Printf("Could not find all users: %s", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h UserRoutes) GetUser(c *gin.Context) {
	var user UserDto
	param := c.Param("user_id")
	user_id, err := strconv.Atoi(param)
	if err != nil {
		log.Printf("Problem with provided user id '%s': %s", param, err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	user, err = h.svc.FindUser(user_id)
	if err != nil {
		log.Printf("Did not find user %d: %s", user_id, err)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, user)
}

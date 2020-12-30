package main

import (
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int
	Name string
}

const (
	port = "8081"
)

var users = map[string]User{
	"1": {
		ID:   1,
		Name: "Alice",
	},
	"2": {
		ID:   2,
		Name: "Bob",
	},
}

func main() {
	router := gin.Default()
	router.GET("/users/:id", getUserByID)
	router.Run(net.JoinHostPort("", port))
}

func getUserByID(c *gin.Context) {
	id := c.Param("id")
	timeoutQuery := c.Query("timeout")
	timeout, err := time.ParseDuration(timeoutQuery)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	time.Sleep(timeout)
	user := users[id]
	c.JSON(http.StatusOK, user)
}

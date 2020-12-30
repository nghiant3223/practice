package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
)

const (
	port            = "8080"
	externalBaseURL = "http://127.0.0.1:8081"
	commandName     = "call_user_service"
)

type User struct {
	ID   int
	Name string
}

var baseUser = User{
	ID:   100,
	Name: "NghiaNT",
}

func main() {
	router := gin.Default()
	router.GET("/login/:id", loginUser)
	router.Run(net.JoinHostPort("", port))
}

func loginUser(c *gin.Context) {
	id := c.Param("id")
	timeout := c.Query("timeout")
	user, err := getUserById(id, timeout)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, user)
}

func getUserById(id string, timeout string) (User, error) {
	var user User
	hystrix.ConfigureCommand(commandName, hystrix.CommandConfig{
		RequestVolumeThreshold: 5,
		SleepWindow:            10 * 1000,
	})
	err := hystrix.Do(commandName, func() error {
		url := externalBaseURL + "/users/" + id + "?timeout=" + timeout
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		err = json.Unmarshal(body, &user)
		if err != nil {
			return err
		}
		return nil
	}, func(err error) error {
		log.Printf("error on fallback: %v\n", err)
		user = baseUser
		return nil
	})
	if err != nil {
		log.Printf("error after do: %v\n", err)
		return User{}, err
	}
	return user, nil
}

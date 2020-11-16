package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var route2Handler = map[string]gin.HandlerFunc{
	"/ping":  pingHandler,
	"/greet": greetHandler,
}

func pingHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "PONG")
}

func greetHandler(ctx *gin.Context) {
	username := ctx.Query("username")
	ctx.String(http.StatusOK, "Hello %s", username)
}

func registerHandler(router *gin.Engine) {
	for route, handler := range route2Handler {
		router.GET(route, handler)
	}
}

func NewServer() *gin.Engine {
	router := gin.Default()
	registerHandler(router)
	return router
}

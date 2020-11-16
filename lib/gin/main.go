package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	UseMiddleware(router, func(ctx *gin.Context) {

	})
	UseMiddleware(router, func(ctx *gin.Context) {

	})
}

func UseMiddleware(r gin.IRouter, fn gin.HandlerFunc) {
	r.Use(fn)
}

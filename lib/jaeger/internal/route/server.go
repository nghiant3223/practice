package route

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nghiant3223/gopractice/jaeger/pkg/tracing"
	"github.com/spf13/cast"
)

type Server interface {
	Listen() error
}

type server struct {
	port int
	db   DB
	app  *gin.Engine
}

func NewServer(port int, db DB) Server {
	app := gin.Default()
	server := &server{
		port: port,
		db:   db,
		app:  app,
	}
	server.register()
	return server
}

func (s *server) Listen() error {
	addr := fmt.Sprintf(":%d", s.port)
	return s.app.Run(addr)
}

func (s *server) register() {
	s.app.GET("/routes/:id", s.getRoute)
	s.app.GET("/routes", s.getRoutes)
}

func (s *server) getRoute(c *gin.Context) {
	_, ctx, finish := tracing.StartSpanFromGinContext(c)
	defer finish()

	paramID := c.Param("id")
	routeID := cast.ToUint64(paramID)
	route, err := s.db.Find(ctx, routeID)
	if err != nil {
		if errors.Is(err, ErrorNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, route)
}

func (s *server) getRoutes(c *gin.Context) {
	_, ctx, finish := tracing.StartSpanFromGinContext(c)
	defer finish()

	routes, err := s.db.FindAll(ctx)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, routes)
}

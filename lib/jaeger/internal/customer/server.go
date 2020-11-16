package customer

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
	s.app.GET("/customers/:id", s.getCustomer)
	s.app.GET("/customers", s.getCustomers)
}

func (s *server) getCustomer(c *gin.Context) {
	_, ctx, finish := tracing.StartSpanFromGinContext(c)
	defer finish()

	paramCustomerID := c.Param("id")
	customerID := cast.ToUint64(paramCustomerID)
	customer, err := s.db.Find(ctx, customerID)
	if err != nil {
		if errors.Is(err, ErrorNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, customer)
}

func (s *server) getCustomers(c *gin.Context) {
	_, ctx, finish := tracing.StartSpanFromGinContext(c)
	defer finish()

	customers, err := s.db.FindAll(ctx)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, customers)
}

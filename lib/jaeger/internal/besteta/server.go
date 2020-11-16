package besteta

import (
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
	port    int
	service Service
	app     *gin.Engine
}

func NewServer(port int, service Service) Server {
	app := gin.Default()
	server := &server{
		port:    port,
		service: service,
		app:     app,
	}
	server.register()
	return server
}

func (s *server) Listen() error {
	addr := fmt.Sprintf(":%d", s.port)
	return s.app.Run(addr)
}

func (s *server) register() {
	s.app.GET("/besteta/:customerID", s.getBestEtaForCustomerByID)
}

func (s *server) getBestEtaForCustomerByID(c *gin.Context) {
	_, ctx, finish := tracing.StartSpanFromGinContext(c)
	defer finish()

	paramCustomerID := c.Param("customerID")
	customerID := cast.ToUint64(paramCustomerID)
	c.Request.Context()
	route, err := s.service.GetBestRouteForCustomer(ctx, customerID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, route)
}

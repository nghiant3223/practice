package besteta

import (
	"context"
	"fmt"
	"math"

	"github.com/nghiant3223/gopractice/jaeger/internal/customer"
	"github.com/nghiant3223/gopractice/jaeger/internal/route"
	"github.com/nghiant3223/gopractice/jaeger/pkg/http"
	"github.com/nghiant3223/gopractice/jaeger/pkg/tracing"
	"github.com/spf13/cast"
)

type Service interface {
	GetBestRouteForCustomer(ctx context.Context, userID uint64) (route.Route, error)
}

type service struct {
	httpClient http.Client
}

func NewService(httpClient http.Client) Service {
	return &service{
		httpClient: httpClient,
	}
}

func (s *service) GetBestRouteForCustomer(ctx context.Context, customerID uint64) (route.Route, error) {
	_, ctx, finish := tracing.StartSpanFromContext(ctx, "BESTETA_GET_BEST_ROUTE_FOR_CUSTOMER")
	defer finish()

	cx, err := s.getCustomer(ctx, customerID)
	if err != nil {
		return route.Route{}, err
	}
	routes, err := s.getAllRoutes(ctx)
	if err != nil {
		return route.Route{}, err
	}
	bestRoute := s.findBestRoute(ctx, cx, routes)
	return bestRoute, nil
}

func (s *service) getCustomer(ctx context.Context, customerID uint64) (customer.Customer, error) {
	_, ctx, finish := tracing.StartSpanFromContext(ctx, "BESTETA_GET_CUSTOMER")
	defer finish()

	url := fmt.Sprintf("http://localhost:8080/customers/%d", customerID)
	var cx customer.Customer
	err := s.httpClient.Get(ctx, url, &cx)
	if err != nil {
		return customer.Customer{}, err
	}
	return cx, nil
}

func (s *service) getAllRoutes(ctx context.Context, ) ([]route.Route, error) {
	_, ctx, finish := tracing.StartSpanFromContext(ctx, "BESTETA_GET_ALL_ROUTES")
	defer finish()

	url := fmt.Sprintf("http://localhost:8081/routes")
	var routes []route.Route
	err := s.httpClient.Get(ctx, url, &routes)
	if err != nil {
		return nil, err
	}
	return routes, nil
}

func (s *service) findBestRoute(ctx context.Context, cx customer.Customer, routes []route.Route) route.Route {
	_, ctx, finish := tracing.StartSpanFromContext(ctx, "BESTETA_FIND_BEST_ROUTE")
	defer finish()

	min := math.MaxFloat64
	var minRoute route.Route
	for _, rt := range routes {
		distance := euclidDistance(cx.X, cx.Y, rt.X, rt.Y)
		if distance < min {
			min = distance
			minRoute = rt
		}
	}
	return minRoute
}

func euclidDistance(x1, y1, x2, y2 int) float64 {
	squareDistance := (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
	return math.Sqrt(cast.ToFloat64(squareDistance))
}

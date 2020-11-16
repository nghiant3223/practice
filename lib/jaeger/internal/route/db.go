package route

import (
	"context"

	"github.com/nghiant3223/gopractice/jaeger/pkg/tracing"
)

type DB interface {
	Find(ctx context.Context, id uint64) (Route, error)
	FindAll(ctx context.Context, ) ([]Route, error)
}

type sqlDB struct{}

func NewDB() DB {
	return &sqlDB{}
}

func (s *sqlDB) Find(ctx context.Context, id uint64) (Route, error) {
	_, ctx, finish := tracing.StartSpanFromContext(ctx, "ROUTE_FIND")
	defer finish()

	for _, route := range collection {
		if route.ID == id {
			return route, nil
		}
	}
	return Route{}, ErrorNotFound
}

func (s *sqlDB) FindAll(ctx context.Context) ([]Route, error) {
	_, ctx, finish := tracing.StartSpanFromContext(ctx, "ROUTE_FIND_ALL")
	defer finish()

	i := 0
	routes := make([]Route, len(collection))
	for _, route := range collection {
		routes[i] = route
		i++
	}
	return routes, nil
}

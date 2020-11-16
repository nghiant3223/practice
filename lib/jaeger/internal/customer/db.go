package customer

import (
	"context"

	"github.com/nghiant3223/gopractice/jaeger/pkg/log"
	"github.com/nghiant3223/gopractice/jaeger/pkg/tracing"
	"go.uber.org/zap"
)

type DB interface {
	Find(ctx context.Context, id uint64) (Customer, error)
	FindAll(ctx context.Context) ([]Customer, error)
}

type sqlDB struct{}

func NewDB() DB {
	return &sqlDB{}
}

func (s *sqlDB) Find(ctx context.Context, id uint64) (Customer, error) {
	_, ctx, finish := tracing.StartSpanFromContext(ctx, "CUSTOMER_FIND")
	defer finish()

	lg := log.For(ctx)

	for _, customer := range table {
		if customer.ID == id {
			lg.Info("customer found", zap.Any("customer", customer))
			return customer, nil
		}
	}

	lg.Error("customer not found", zap.Uint64("user_id", id))
	return Customer{}, ErrorNotFound
}

func (s *sqlDB) FindAll(ctx context.Context) ([]Customer, error) {
	_, ctx, finish := tracing.StartSpanFromContext(ctx, "CUSTOMER_FIND_ALL")
	defer finish()

	return table, nil
}

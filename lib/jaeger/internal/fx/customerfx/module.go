package customerfx

import (
	"context"

	"github.com/nghiant3223/gopractice/jaeger/internal/customer"
	"github.com/nghiant3223/gopractice/jaeger/pkg/log"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Options(
	fx.Invoke(initServer),
	fx.Provide(provideDB),
)

func initServer(lc fx.Lifecycle, db customer.DB) customer.Server {
	port := 8080
	server := customer.NewServer(port, db)
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				err := server.Listen()
				if err != nil {
					log.Error("cannot start besteta server", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("customer server stop")
			return nil
		},
	})
	return server
}

func provideDB() customer.DB {
	return customer.NewDB()
}

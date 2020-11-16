package bestetafx

import (
	"context"

	"github.com/nghiant3223/gopractice/jaeger/internal/besteta"
	"github.com/nghiant3223/gopractice/jaeger/pkg/http"
	"github.com/nghiant3223/gopractice/jaeger/pkg/log"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Options(
	fx.Invoke(initServer),
	fx.Provide(provideService),
)

func initServer(lc fx.Lifecycle, service besteta.Service) besteta.Server {
	port := 8082
	server := besteta.NewServer(port, service)
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
			log.Info("besteta server stop")
			return nil
		},
	})
	return server
}

func provideService(client http.Client) besteta.Service {
	return besteta.NewService(client)
}

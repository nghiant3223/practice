package tracingfx

import (
	"context"

	"github.com/nghiant3223/gopractice/jaeger/pkg/tracing"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/fx"
)

func InitializeTracer(serviceName string) func(lifecycle fx.Lifecycle) (opentracing.Tracer, error) {
	return func(lc fx.Lifecycle) (opentracing.Tracer, error) {
		tracer, closeFn, err := tracing.Init(serviceName)
		if err != nil {
			return nil, err
		}
		lc.Append(fx.Hook{
			OnStop: func(ctx context.Context) error {
				return closeFn()
			},
		})
		return tracer, nil
	}
}

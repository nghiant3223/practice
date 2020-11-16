package main

import (
	"github.com/nghiant3223/gopractice/jaeger/internal/fx/customerfx"
	"github.com/nghiant3223/gopractice/jaeger/internal/fx/tracingfx"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Invoke(tracingfx.InitializeTracer("customer")),
		customerfx.Module,
	).Run()
}

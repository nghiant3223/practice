package main

import (
	"github.com/nghiant3223/gopractice/jaeger/internal/fx/bestetafx"
	"github.com/nghiant3223/gopractice/jaeger/internal/fx/tracingfx"
	"github.com/nghiant3223/gopractice/jaeger/pkg/http"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Invoke(tracingfx.InitializeTracer("besteta")),
		bestetafx.Module,
		http.Module,
	).Run()
}

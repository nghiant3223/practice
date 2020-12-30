package main

import (
	"context"
	"log"
	"time"

	"go.uber.org/fx"
)

type foo struct {
}

func provideRedis(lc fx.Lifecycle) foo {
	f := foo{}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Println("on")
			return nil
		},
	})
	return f
}

func start(f foo) {
	log.Println(f)
	log.Println("start server")
	time.Sleep(3*time.Second)
}

func main() {
	fx.New(
		fx.Provide(provideRedis),
		fx.Invoke(start),
	).Run()
}

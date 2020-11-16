package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/nghiant3223/gopractice/grpcgw/internal/handlers"

	gw "github.com/nghiant3223/gopractice/grpcgw/protogens/protos"

	"google.golang.org/grpc"
)

func main() {
	go func() {
		listener, err := net.Listen("tcp", ":9090")
		if err != nil {
			log.Fatal("fail to start server on port 9090")
		}
		grpcServer := grpc.NewServer()
		gw.RegisterBookServiceServer(grpcServer, &handlers.BookServiceHandler{})
		err = grpcServer.Serve(listener)
		if err != nil {
			log.Fatal("cannot start grpc server on port 9090")
		}
		log.Println("grpc server listening on 9090")
	}()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterBookServiceHandlerFromEndpoint(ctx, mux, "localhost:9090", opts)
	if err != nil {
		log.Fatal(err)
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Println(err)
	}
}

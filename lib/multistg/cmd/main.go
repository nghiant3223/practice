package main

import (
	"log"

	"github.com/nghiant3223/gopractice/multistg/internal"
)

func main() {
	server := internal.NewServer()
	err := server.Run()
	if err != nil {
		log.Printf("failed to start server; %v", err)
	}
}

package main

import (
	"net"
	"net/http"
)

func main() {
	http.ListenAndServe(net.JoinHostPort("", "8080"), nil)
}

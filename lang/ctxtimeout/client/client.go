package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func makeRequest(ctx context.Context) (string, error) {
	log.Print("make request")
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)
	if err != nil {
		return "", err
	}
	cl := &http.Client{}
	okCh := make(chan string)
	errCh := make(chan error)
	go func() {
		res, err := cl.Do(req)
		if err != nil {
			log.Println("error while making request")
			errCh <- err
			return
		}
		bytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			errCh <- err
			return
		}
		okCh <- string(bytes)
	}()
	select {
	case err := <-errCh:
		return "", err
	case body := <-okCh:
		return body, nil
	case <-ctx.Done():
		select {
		case <-errCh:
			return "", ctx.Err()
		case <-okCh:
			return "", ctx.Err()
		}
	}
}

func main() {
	bgCtx := context.Background()
	timeout := 2 * time.Second
	ctx, cancel := context.WithTimeout(bgCtx, timeout)
	defer cancel()
	body, err := makeRequest(ctx)
	if err != nil {
		log.Print(err)
		return
	}
	log.Print(body)
}

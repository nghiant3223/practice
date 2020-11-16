package main

import (
	"fmt"
	"time"
)

const (
	delay = 100 * time.Millisecond
)

func main() {
	calcFib(45)
}

func calcFib(n int) {
	doneC := make(chan struct{})
	go spin(doneC)
	result := fib(n)
	doneC <- struct{}{}
	fmt.Printf("\nfib(%d)=%d\n", n, result)
}

func spin(doneC chan struct{}) {
	for {
		for _, r := range `-/|\` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
			select {
			case <-doneC:
				return
			default:
			}
		}
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

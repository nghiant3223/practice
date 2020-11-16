package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := listenCancelEvent()
	countdown(abort)
}

func countdown(abort chan struct{}) {
	tick := time.NewTicker(time.Second)
	for i := 10; i > 0; i-- {
		select {
		case <-tick.C:
			fmt.Printf("Rocket gonna be lauched in %d\n", i)
		case <-abort:
			fmt.Println("Launch canceled")
			tick.Stop()
			return
		}
	}
}

func listenCancelEvent() chan struct{} {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(abort)
	}()
	return abort
}

package main

import "fmt"

func main() {
	var fns []func()
	for i := 0; i < 10; i++ {
		fns = append(fns, func() {
			fmt.Println(i)
		})
	}
	for _, fn := range fns {
		fn()
	}
}

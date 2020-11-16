package main

import "fmt"

func foo(s []int) {
	s = append(s, 0)
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(cap(s), len(s))
	s = s[3:]
	fmt.Println(cap(s), len(s))
	//foo(s)
	//fmt.Println(s)
}

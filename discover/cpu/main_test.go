package main

import (
	"fmt"
	"runtime"
	"testing"
)

func BenchmarkSequential(b *testing.B) {
	fmt.Println(runtime.NumCPU())
	for i := 0; i < b.N; i++ {
		add(numbers)
	}
}

func BenchmarkConcurrent(b *testing.B) {
	fmt.Println(runtime.NumCPU())
	for i := 0; i < b.N; i++ {
		addConcurrent(runtime.NumCPU(), numbers)
	}
}

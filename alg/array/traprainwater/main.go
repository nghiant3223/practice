package main

import (
	"fmt"
	"math"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func sumBetweenIAndJ(bars []int, i, j int) int {
	s := 0
	for k := i + 1; k < j; k++ {
		s += bars[k]
	}
	return s
}

func indexOfMax(bars []int, i, j int) int {
	var l int
	m := math.MinInt64
	for k := i; k <= j; k++ {
		c := bars[k]
		if c > m {
			m = c
			l = k
		}
	}
	return l
}

func calcTrapArea(bars []int, i, j int) int {
	width := j - i - 1
	height := min(bars[i], bars[j])
	return width*height - sumBetweenIAndJ(bars, i, j)
}

func moveLeft(bars []int, i, pivot int) int {
	if i <= 0 {
		return 0
	}
	indexOfMax := indexOfMax(bars, 0, i)
	trapArea := calcTrapArea(bars, indexOfMax, pivot)
	return trapArea + moveLeft(bars, indexOfMax-1, indexOfMax)
}

func moveRight(bars []int, i, pivot int) int {
	length := len(bars)
	if i >= length {
		return 0
	}
	indexOfMax := indexOfMax(bars, i, length-1)
	trapArea := calcTrapArea(bars, pivot, indexOfMax)
	return trapArea + moveRight(bars, indexOfMax+1, indexOfMax)
}

func trap(bars []int) int {
	i := indexOfMax(bars, 0, len(bars)-1)
	return moveLeft(bars, i-1, i) + moveRight(bars, i+1, i)
}

func main() {
	bars := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	ans := trap(bars)
	fmt.Print(ans)
}

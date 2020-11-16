package main

import (
	"fmt"
	"math"
)

func secondMax(nums []int) int {
	max := math.MinInt32
	secondMax := math.MinInt32
	for _, n := range nums {
		if n > max {
			secondMax = max
			max = n
			continue
		}
		if n < max && n > secondMax {
			secondMax = n
			continue
		}
	}
	return secondMax
}

func main() {
	nums := []int{9, 5, 7, 3, 20}
	ans := secondMax(nums)
	fmt.Print(ans)
}

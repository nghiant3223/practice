package main

import (
	"fmt"
	"sort"
)

func lastStoneWeight(stones []int) int {
	sort.Ints(stones)
	i := len(stones) - 1
	for ; i > 1; i-- {
		if stones[i] > stones[i-1] {
			stones[i-1] = stones[i] - stones[i-1]
			continue
		}
		i--
	}
	if i < 0 {
		return 0
	}
	return stones[0]
}

func main() {
	stones := []int{2, 7, 4, 1, 8, 1}
	ans := lastStoneWeight(stones)
	fmt.Println(ans)
}

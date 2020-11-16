package main

import (
	"fmt"
)

var solution [][]int

func combinationSum(candidates []int, target int) {
	combinationSumUtil(candidates, target, []int{}, 0)
}

func combinationSumUtil(candidates []int, target int, combination []int, i int) {
	if target < 0 {
		return
	}
	if target == 0 {
		solution = append(solution, combination)
		return
	}
	l := len(candidates)
	for j := i; j < l; j++ {
		nextCombination := append(combination, candidates[j])
		combinationSumUtil(candidates, target-candidates[j], nextCombination, j)
	}
}

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	combinationSum(candidates, target)
	fmt.Println(solution)
}

package main

import "fmt"

// Given a set of candidate numbers (candidates) (without duplicates) and a target number (target),
// find all unique combinations in candidates where the candidate numbers sums to target.

func combinationSum(candidates []int, target int) [][]int {
	var solution [][]int
	backtrack(candidates, &solution, []int{}, 0, 0, target)
	return solution
}

func backtrack(candidates []int, solutions *[][]int, subSolution []int, idx int, sum int, target int) {
	if sum == target {
		*solutions = append(*solutions, subSolution)
		return
	}
	if sum > target {
		return
	}
	for i := idx; i < len(candidates); i++ {
		cdd := candidates[i]
		tmpSolution := append(subSolution, cdd)
		backtrack(candidates, solutions, tmpSolution, i+1, sum+cdd, target)
	}
}

func main() {
	candidates := []int{2, 4, 6, 10}
	target := 16
	solution := combinationSum(candidates, target)
	fmt.Println(solution)
}

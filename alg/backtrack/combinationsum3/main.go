package main

import "fmt"

// Find all possible combinations of k numbers that add up to a number n,
// given that only numbers from 1 to 9 can be used and each combination should be a unique set of numbers.

func combinationSum3(k int, n int) [][]int {
	var solutions [][]int
	backtrack([]int{}, 0, k, 0, n, &solutions)
	return solutions
}

func backtrack(items []int, i int, k int, sum int, n int, solutions *[][]int) {
	if len(items) > k {
		return
	}
	if sum > n {
		return
	}
	if len(items) == k && sum == n {
		*solutions = append(*solutions, items)
		return
	}
	for j := i + 1; j < n; j++ {
		arr := make([]int, len(items))
		copy(arr, items)
		backtrack(append(arr, j), j, k, sum+j, n, solutions)
	}
	return
}

func main() {
	result := combinationSum3(3, 9)
	fmt.Println(result)
}

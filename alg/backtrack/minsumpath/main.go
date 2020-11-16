package main

import (
	"fmt"
	"math"
)

// Given a m x n grid filled with non-negative numbers,
// find a path from top left to bottom right which minimizes the sum of all numbers along its path.

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minPathSum(grid [][]int) int {
	ans := math.MaxInt64
	findMinPathSum(grid, 0, 0, grid[0][0], &ans)
	return ans
}

func findMinPathSum(grid [][]int, r, c int, sum int, bestSoFar *int) {
	m := len(grid)
	n := len(grid[0])
	// Position is out of range
	if r == m || c == n {
		return
	}
	// If bottom-right cell is reached, get the best so far
	if r == m-1 && c == n-1 {
		*bestSoFar = min(sum, *bestSoFar)
		return
	}
	// If there is a solution and current sum is greater than best solution so far
	// We're not gonna progress this temporary solution
	if bestSoFar != nil && sum >= *bestSoFar {
		return
	}
	// Go right
	findMinPathSum(grid, r, c+1, sum+grid[r][c], bestSoFar)
	// Go down
	findMinPathSum(grid, r+1, c, sum+grid[r][c], bestSoFar)
}

func main() {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	ans := minPathSum(grid)
	fmt.Println(ans == 7)
}

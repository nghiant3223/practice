package main

import "fmt"

// Problem:
// Given a 2d grid map of '1's (land) and '0's (water), count the number of islands.
// An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
// You may assume all four edges of the grid are all surrounded by water.

// Idea:
// For each land cell, if cell is not visited yet then it increments island count and
// try to expand territory as wide as possible using DFS

var visited map[[2]int]bool

func shouldVisit(grid [][]bool, m, n, i, j int) bool {
	return i >= 0 && j >= 0 && i < m && j < n && grid[i][j] && !visited[[2]int{i, j}]
}

func dfs(grid [][]bool, m, n, i, j int) {
	visited[[2]int{i, j}] = true
	if shouldVisit(grid, m, n, i-1, j) {
		dfs(grid, m, n, i-1, j)
	}
	if shouldVisit(grid, m, n, i+1, j) {
		dfs(grid, m, n, i+1, j)
	}
	if shouldVisit(grid, m, n, i, j-1) {
		dfs(grid, m, n, i, j-1)
	}
	if shouldVisit(grid, m, n, i, j+1) {
		dfs(grid, m, n, i, j+1)
	}
}

func numberOfIslands(grid [][]bool) int {
	cnt := 0
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if visited[[2]int{i, j}] || !grid[i][j] {
				continue
			}
			cnt++
			dfs(grid, m, n, i, j)
		}
	}
	return cnt
}

func main() {
	visited = make(map[[2]int]bool)
	grid := [][]bool{
		{true, true, false, false, false},
		{true, true, false, false, false},
		{false, false, true, false, false},
		{false, false, false, true, true},
	}
	ans := numberOfIslands(grid)
	fmt.Println(ans)
}

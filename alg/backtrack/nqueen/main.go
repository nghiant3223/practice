package main

import "fmt"

const n = 4

func backtrack(board [][]bool, c int) bool {
	if c >= n {
		return true
	}
	for r := 0; r < n; r++ {
		board[r][c] = true
		if isSafe(board, r, c) {
			ok := backtrack(board, c+1)
			if ok {
				return true
			}
		}
		board[r][c] = false
	}
	return false
}

func isSafe(board [][]bool, r, c int) bool {
	for i := c - 1; i >= 0; i-- {
		if board[r][i] {
			return false
		}
	}
	for i, j := r-1, c-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] {
			return false
		}
	}
	return true
}

func initBoard() [][]bool {
	board := make([][]bool, n)
	for i := 0; i < n; i++ {
		board[i] = make([]bool, n)
	}
	return board
}

func printBoard(board [][]bool) {
	for _, r := range board {
		fmt.Println(r)
	}
}

func main() {
	board := initBoard()
	firstCol := 0
	backtrack(board, firstCol)
	printBoard(board)
}

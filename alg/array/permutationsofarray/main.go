package main

import "log"

func main() {
	ans := permutationsOfArray([]int{0, 1, 2, 3, 4})
	log.Println(ans)
}

var solution [][]int

func permutationsOfArray(a []int) [][]int {
	used := make([]bool, len(a))
	permutationsOfArrayUtil(a, []int{}, used)
	return solution
}

func permutationsOfArrayUtil(a []int, cur []int, used []bool) {
	if len(cur) == len(a) {
		solution = append(solution, cur)
		return
	}
	for _, v := range a {
		if used[v] {
			continue
		}
		used[v] = true
		permutationsOfArrayUtil(a, append(cur, v), used)
		used[v] = false
	}
}

package main

import "log"

func main() {
	ans := permutationsOfArray([]int{1, 2, 3})
	log.Println(ans)
}

var solution [][]int

func permutationsOfArray(a []int) [][]int {
	permutationsOfArrayUtil(a, []int{})
	return solution
}

func permutationsOfArrayUtil(a []int, cur []int) {
	if len(cur) == len(a) {
		solution = append(solution, cur)
		return
	}
	for _, v := range a {
		if contains(cur, v) {
			continue
		}
		permutationsOfArrayUtil(a, append(cur, v))
	}
}

func contains(a []int, i int) bool {
	for _, v := range a {
		if v == i {
			return true
		}
	}
	return false
}

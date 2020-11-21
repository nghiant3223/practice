package subsetofarray

var solutions [][]int

func subsetOfArrays(a []int) [][]int {
	subsetOfArrayUtil(a, 0, []int{})
	return solutions
}

func subsetOfArrayUtil(a []int, i int, cur []int) {
	for j := i; j < len(a); j++ {
		tmp := append(cur, a[j])
		solutions = append(solutions, tmp)
		subsetOfArrayUtil(a, j+1, tmp)
	}
}

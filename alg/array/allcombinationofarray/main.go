package allcombinationofarray

func allCombinationOfArray(nums []int) [][]int {
	return allCombinationOfArrayUtil(nums, -1)
}

func allCombinationOfArrayUtil(nums []int, i int) [][]int {
	if i >= len(nums)-1 {
		return [][]int{{}}
	}
	var partialSolution [][]int
	for j := i + 1; j < len(nums); j++ {
		solutions := allCombinationOfArrayUtil(nums, j)
		for _, solution := range solutions {
			solution = append([]int{nums[j]}, solution...)
			partialSolution = append(partialSolution, solution)
		}
	}
	return partialSolution
}

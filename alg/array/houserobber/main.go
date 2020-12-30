package houserobber

func rob(nums []int) int {
	maxRob := 0
	l := len(nums)
	for i := l - 1; i >= 0; i-- {
		if i+2 < l {
			maxI := 0
			for j := i + 2; j < l; j++ {
				maxI = max(maxI, nums[j])
			}
			nums[i] = maxI + nums[i]
		}
		maxRob = max(maxRob, nums[i])
	}
	return maxRob
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package bestsumtarget

// Given an target and array of numbers
// Find the smallest combination of numbers that can add up to target
// Numbers can be reused as needed

func bestSum(target int, numbers []int) []int {
	dp := make([][]int, target+1)
	for i := 1; i <= target; i++ {
		var smallestCombination []int
		for _, number := range numbers {
			candidate := i - number
			if candidate < 0 {
				continue
			}
			if candidate == 0 {
				dp[i] = []int{number}
				smallestCombination = dp[i]
				break
			}
			if len(dp[candidate]) == 0 {
				continue
			}
			if len(smallestCombination) == 0 || len(dp[candidate])+1 < len(smallestCombination) {
				smallestCombination = append(dp[candidate], number)
				continue
			}
		}
		dp[i] = smallestCombination
	}
	return dp[target]
}

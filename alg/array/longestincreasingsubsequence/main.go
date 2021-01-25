package longestincreasingsubsequence

import "math"

func lengthOfLIS(nums []int) int {
	count := len(nums)
	dp := make([]int, count)
	dp[count-1] = 1

	for i := count - 2; i >= 0; i-- {
		max := 0
		for j := i + 1; j < count; j++ {
			if nums[i] >= nums[j] {
				continue
			}
			max = maxInt(max, dp[j])
		}
		dp[i] = 1 + max
	}

	return maxIntSlice(dp)
}

func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxIntSlice(s []int) int {
	max := math.MinInt64
	for _, n := range s {
		if n > max {
			max = n
		}
	}
	return max
}

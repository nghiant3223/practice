package minnumberofcoins

import "math"

var coins = []int{1, 2, 5}

const (
	infeasible = -1
	maxCoin    = 5
)

//func minNumberOfCoins(target int) int {
//	if target < 0 {
//		return infeasible
//	}
//	if target == 0 {
//		return 0
//	}
//	ans := math.MaxInt64
//	for _, coin := range coins {
//		candidateAns := minNumberOfCoins(target-coin)
//		if candidateAns == infeasible {
//			continue
//		}
//		ans = min(ans, candidateAns+1)
//	}
//	return ans
//}

func minNumberOfCoins(target int) int {
	dp := make([]int, max(maxCoin, target)+1)
	for i := range dp {
		dp[i] = math.MaxInt64
	}
	for _, coin := range coins {
		dp[coin] = 1
	}
	for i := 0; i <= target; i++ {
		cand := math.MaxInt64
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			cand = min(cand, dp[i-coin])
		}
		if cand != math.MaxInt64 {
			dp[i] = min(dp[i], cand+1)
		}
	}
	return dp[target]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

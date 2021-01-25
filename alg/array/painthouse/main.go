package painthouse

func paintHouse(costs [][3]int) int {
	count := len(costs)
	dp := make([][3]int, len(costs))

	dp[count-1][0] = costs[count-1][0]
	dp[count-1][1] = costs[count-1][1]
	dp[count-1][2] = costs[count-1][2]

	for i := count - 2; i >= 0; i-- {
		dp[i][0] = costs[i][0] + min(dp[i+1][1], dp[i+1][2])
		dp[i][1] = costs[i][1] + min(dp[i+1][0], dp[i+1][2])
		dp[i][2] = costs[i][2] + min(dp[i+1][0], dp[i+1][1])
	}

	return min(dp[0][0], min(dp[0][1], dp[0][2]))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

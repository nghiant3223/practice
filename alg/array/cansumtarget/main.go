package cansumtarget

func canSum(target int, numbers []int) bool {
	dp := make([]bool, target+1)
	dp[0] = true

	for i := 1; i < target+1; i++ {
		for _, number := range numbers {
			candidateNumber := i - number
			if candidateNumber < 0 {
				continue
			}
			if dp[candidateNumber] {
				dp[i] = true
				break
			}
		}
	}

	return dp[target]
}

package decodeways

func numDecodings(s string) int {
	c := len(s)
	dp := make([]int, c)
	if s[c-1] != '0' {
		dp[c-1] = 1
	} else {
		dp[c-1] = 0
	}
	if s[c-2] == '0' || s[c-2:] == "00" {
		dp[c-2] = 0
	} else if s[c-2:] == "10" || s[c-2:] == "20" {
		dp[c-2] = 1
	} else if s[c-2:] <= "27" {
		dp[c-2] = 2
	} else {
		dp[c-2] = 1
	}
	for i := c - 3; i >= 0; i-- {
		char := s[i]
		switch char {
		case '0':
			dp[i] = 0
		case '1':
			if s[i+1] == '0' {
				dp[i] = dp[i+2]
			} else {
				dp[i] = dp[i+1] + dp[i+2]
			}
		case '2':
			if s[i+1] == '0' {
				dp[i] = dp[i+2]
			} else if s[i+1] <= '7' {
				dp[i] = dp[i+1] + dp[i+2]
			} else {
				dp[i] = dp[i+1]
			}
		default:
			dp[i] = dp[i+1]
		}
	}
	return dp[0]
}

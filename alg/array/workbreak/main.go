package workbreak

import "strings"

// Given an target string and a word bank, check whether the target word can be constructed from words in the word bank
// You may reuse elements of words in bank as many times as needed

func wordBreak(s string, wordDict []string) bool {
	if s == "" {
		return true
	}
	for _, word := range wordDict {
		if !strings.HasPrefix(s, word) {
			continue
		}
		if wordBreak(s[len(word):], wordDict) {
			return true
		}
	}
	return false
}

func wordBreakDP(s string, wordDict []string) bool {
	count := len(s)
	dp := make([]bool, count+1)
	dp[0] = true
	for i := range s {
		if !dp[i] {
			continue
		}
		for j := i + 1; j <= count; j++ {
			if contains(wordDict, s[i:j]) {
				dp[j] = true
			}
		}
	}
	return dp[count]
}

func contains(ss []string, s string) bool {
	for _, candidate := range ss {
		if s == candidate {
			return true
		}
	}
	return false
}

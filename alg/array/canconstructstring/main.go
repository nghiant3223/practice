package canconstructstring

import "strings"

// Given an target string and a word bank, check whether the target word can be constructed from words in the word bank
// You may reuse elements of words in bank as many times as needed

func canConstruct(target string, banks []string) bool {
	if target == "" {
		return true
	}
	for _, word := range banks {
		if !strings.HasPrefix(target, word) {
			continue
		}
		if canConstruct(target[len(word):], banks) {
			return true
		}
	}
	return false
}
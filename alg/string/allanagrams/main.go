package allanagrams

import "strings"

// Input: S = "abacebca", T = "abc"
// Output: [0,1,5]

func allAnagrams(s string, t string) []int {
	var answers []int
	sCount := len(s)
	tCount := len(t)
	b, e := 0, 1
	for e < sCount {
		if isCandidate(t, s[e]) {
			if e-b+1 < tCount {
				e++
			} else {
				if isAnagram(s[b:e+1], t) {
					answers = append(answers, b)
				}
				b++
				e++
			}
		} else {
			b = e
			e++
		}
	}
	return answers
}

func isCandidate(t string, char byte) bool {
	return strings.IndexByte(t, char) > -1
}

func isAnagram(a, b string) bool {
	aF := make([]int, 26)
	bF := make([]int, 26)
	for _, c := range a {
		aF[c-'a']++
	}
	for _, c := range b {
		bF[c-'a']++
	}
	for i := range aF {
		if aF[i] != bF[i] {
			return false
		}
	}
	return true
}

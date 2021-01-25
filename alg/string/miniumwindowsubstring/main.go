package miniumwindowsubstring

// Given a string S and a string T, find the minimum window in S which will contain all the characters in T in complexity O(N)
//
// Example:
// Input: S = "ADOBECODEBANC", T = "ABC"
// Output: "BANC"

// Note:
// - If there is no s uch window in S that covers all characters in T, return the empty string ""
// - If there is such window, you are guaranteed that there will always be only one unique minimum window in S

func minimumWindowSubstring(s, t string) string {
	b, e := 0, 1
	sLen := len(s)
	for e < sLen {
		if !sContainsT(s[b:e], t) {
			e++
			continue
		}
		b++
	}
	tmp := s[b:e]
	var prev string
	for sContainsT(tmp, t) {
		prev = tmp
		tmp = tmp[1:]
	}
	return prev
}

func sContainsT(s, t string) bool {
	lookup := make(map[byte]struct{})
	for i := range s {
		lookup[s[i]] = struct{}{}
	}
	for i := range t {
		_, ok := lookup[t[i]]
		if !ok {
			return false
		}
	}
	return true
}

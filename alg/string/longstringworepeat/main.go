package main

import (
	"fmt"
	"strings"
)

// Problem: Given a string s, find the length of the longest substring without repeating characters.

func lengthOfLongestSubstring(s string) int {
	b := 0
	e := 0
	tmp := ""
	ans := 0
	n := len(s)
	for e < n {
		ch := string(s[e])
		i, found := search(tmp, ch)
		if !found {
			tmp += ch
			ans = max(ans, e-b+1)
			e++
		} else {
			b = i + b + 1
			tmp = tmp[i+1:]
		}
	}
	return max(ans, b-e+1)
}

func search(s string, r string) (int, bool) {
	i := strings.Index(s, r)
	if i >= 0 {
		return i, true
	}
	return i, false
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	ans := lengthOfLongestSubstring("pwwkew")
	fmt.Println(ans)
}

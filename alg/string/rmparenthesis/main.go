package main

import (
	"fmt"

	"github.com/nghiant3223/gopractice/alg/string/ds"
)

func minRemoveToMakeValidParentheses(s string) string {
	st := ds.Stack{}
	var indicesToRm []int
	for i, c := range s {
		if c == '(' {
			st.Push(i)
			continue
		}
		if c == ')' {
			if !st.IsEmpty() {
				st.Pop()
				continue
			}
			indicesToRm = append(indicesToRm, i)
		}
	}
	for !st.IsEmpty() {
		top := st.Pop()
		indicesToRm = append(indicesToRm, top.(int))
	}
	result := s
	for _, i := range indicesToRm {
		result = removeAt(result, i)
	}
	return result
}

func removeAt(s string, i int) string {
	return s[:i] + s[i+1:]
}

func main() {
	fmt.Println(minRemoveToMakeValidParentheses("(a(b(c)d)"))
}

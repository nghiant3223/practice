package main

import (
	"fmt"
	"strings"

	"github.com/nghiant3223/gopractice/alg/string/ds"
)

func simplifyPath(path string) string {
	pathComponents := strings.Split(path, "/")
	st := ds.Stack{}
	for _, c := range pathComponents {
		if c == "." || c == "" {
			continue
		}
		if c == ".." {
			if st.IsEmpty() {
				continue
			}
			st.Pop()
			continue
		}
		st.Push(c)
	}
	resultComponents := make([]string, st.Count())
	i := st.Count() - 1
	for !st.IsEmpty() {
		resultComponents[i] = st.Pop()
		i--
	}
	return "/" + strings.Join(resultComponents, "/")
}

func main() {
	fmt.Println(simplifyPath("/home/"))
	fmt.Println(simplifyPath("/../"))
	fmt.Println(simplifyPath("/home//foo/"))
	fmt.Println(simplifyPath("/a/./b/../../c/"))
	fmt.Println(simplifyPath("/a/../../b/../c//.//"))
	fmt.Println(simplifyPath("/a//b////c/d//././/.."))
}

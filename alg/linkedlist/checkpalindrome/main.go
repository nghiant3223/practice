package checkpalindrome

import (
	"strconv"

	"github.com/nghiant3223/gopractice/alg/linkedlist/ds"
)

func isPalindrome(head *ds.Node) bool {
	s1 := traverse(head)
	s2 := reverseTraverse(head)
	return s1 == s2
}

func traverse(node *ds.Node) string {
	if node == nil {
		return ""
	}
	return strconv.Itoa(node.Value) + traverse(node.Next)
}

func reverseTraverse(node *ds.Node) string {
	if node == nil {
		return ""
	}
	return reverseTraverse(node.Next) + strconv.Itoa(node.Value)
}

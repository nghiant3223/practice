package util

import (
	"fmt"

	"github.com/nghiant3223/gopractice/alg/linkedlist/ds"
)

func PrintLinkedList(root *ds.Node) {
	head := root
	for head != nil {
		fmt.Printf("%v->", head.Value)
		head = head.Next
	}
	fmt.Print("NULL\n")
}

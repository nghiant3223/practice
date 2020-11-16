package out

import (
	"log"
	"strings"

	"github.com/nghiant3223/gopractice/alg/tree/ds"
)

const (
	indentCharacter = " "
)

func Print(node *ds.Node, depth ...int) {
	if node == nil {
		return
	}
	realDepth := 0
	if len(depth) != 0 {
		realDepth = depth[0]
	}
	indent := strings.Repeat(indentCharacter, realDepth)
	log.Printf("%s%d\n", indent, node.Value)
	Print(node.Left, []int{realDepth + 1}...)
	Print(node.Right, []int{realDepth + 1}...)
}

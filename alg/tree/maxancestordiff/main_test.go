package maxancestordiff

import (
	"fmt"
	"testing"

	"github.com/nghiant3223/gopractice/alg/tree/ds"
	"github.com/stretchr/testify/assert"
)

func TestHappy(t *testing.T) {
	for i, tc := range []struct {
		setup func() *ds.Node
		out   int
	}{
		{
			setup: func() *ds.Node {
				node8 := ds.NewNode(8)
				node3 := ds.NewNode(3)
				node1 := ds.NewNode(1)
				node6 := ds.NewNode(6)
				node4 := ds.NewNode(4)
				node7 := ds.NewNode(7)
				node10 := ds.NewNode(10)
				node14 := ds.NewNode(14)
				node13 := ds.NewNode(13)
				node8.Left = node3
				node8.Right = node10
				node3.Left = node1
				node3.Right = node6
				node10.Left = node14
				node14.Left = node13
				node6.Left = node4
				node6.Right = node7
				return node8
			},
			out: 7,
		},
		{
			setup: func() *ds.Node {
				node1 := ds.NewNode(1)
				return node1
			},
			out: 0,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			root := tc.setup()
			ans := maxAncestorDiff(root)
			assert.Equal(t, tc.out, ans)
		})
	}
}

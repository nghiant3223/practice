package minimumdepth

import (
	"testing"

	"github.com/nghiant3223/gopractice/alg/tree/ds"
	"github.com/stretchr/testify/assert"
)

func TestHappy(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   *ds.Node
		out  int
	}{
		{
			name: "1",
			in:   ds.ParseTree([]int{8, 3, 10, 1, 6, 14, 7, ds.Nil, 3, 4, 9, 13, ds.Nil, ds.Nil, ds.Nil, ds.Nil}),
			out:  3,
		},
		{
			name: "2",
			in:   ds.ParseTree([]int{8, 3, 10, 1, 6, ds.Nil, ds.Nil, ds.Nil, 3, 4, 9, ds.Nil, ds.Nil, ds.Nil, ds.Nil, ds.Nil}),
			out:  2,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := minimumDepth(tc.in)
			assert.Equal(t, tc.out, out)
		})
	}
}

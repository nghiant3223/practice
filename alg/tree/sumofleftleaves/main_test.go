package sumofleftleaves

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
			in:   ds.ParseTree([]int{8, 3, 10, 1, 6, 14, ds.Nil, ds.Nil, ds.Nil, 4, 7, 13, ds.Nil, ds.Nil, ds.Nil}),
			out:  18,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := sumOfLeftLeaves(tc.in)
			assert.Equal(t, tc.out, out)
		})
	}
}

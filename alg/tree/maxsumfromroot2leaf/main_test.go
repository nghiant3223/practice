package maxsumfromroot2leaf

import (
	"testing"

	"github.com/nghiant3223/gopractice/alg/tree/ds"
	"github.com/stretchr/testify/assert"
)

func TestHappy(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   *ds.Node
		out1 int
		out2 []int
	}{
		{
			name: "1",
			in:   ds.ParseTree([]int{1, 2, 3, 8, 4, 5, 6, ds.Nil, ds.Nil, 10, ds.Nil, 7, 9, ds.Nil, 5}),
			out1: 18,
			out2: []int{1, 3, 5, 9},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			sum, path := maxSumFromRoot2LeafPath(tc.in)
			assert.Equal(t, tc.out1, sum)
			assert.Equal(t, tc.out2, path)
			m := make(map[int]int, 2)
			m[0] = 1
			m[1] = 2
			assert.Len(t, m, 2)
		})
	}
}

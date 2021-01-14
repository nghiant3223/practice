package isvalidbst

import (
	"testing"

	"github.com/nghiant3223/gopractice/alg/tree/ds"
	"github.com/stretchr/testify/assert"
)

func TestHappy(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   []int
		out  bool
	}{
		{
			name: "1",
			in:   []int{6, 2, 9, 1, 3, 7, 10, ds.Nil, ds.Nil, ds.Nil, ds.Nil, 4, ds.Nil, ds.Nil, ds.Nil},
			out:  false,
		},
		{
			name: "2",
			in:   []int{6, 2, 9, 1, 3, 7, 10},
			out:  true,
		},
		{
			name: "3",
			in:   []int{6, 2, 9, 3, 1, 7, 10},
			out:  false,
		},
		{
			name: "4",
			in:   []int{6, 2, 9, 1, 3, ds.Nil, ds.Nil},
			out:  true,
		},
		{
			name: "1",
			in:   []int{6, 2, 9, 1, 3, 8, 10, ds.Nil, ds.Nil, ds.Nil, ds.Nil, 7, ds.Nil, ds.Nil, ds.Nil},
			out:  true,
		},
		{
			name: "1",
			in:   []int{6, 2, 9, 1, 13, 7, 10, ds.Nil, ds.Nil, ds.Nil, ds.Nil, 4, ds.Nil, ds.Nil, ds.Nil},
			out:  false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			prev = nil
			in := ds.ParseTree(tc.in)
			out := isValidBST(in)
			assert.Equal(t, tc.out, out)
		})
	}
}

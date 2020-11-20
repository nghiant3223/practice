package choptree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinChopCount(t *testing.T) {
	type testcase struct {
		name string
		in   []int
		exp  int
	}
	for _, tc := range []testcase{
		{
			name: "ever-decreasing",
			in:   []int{8, 7, 6, 5},
			exp:  1,
		},
		{
			name: "ever-decreasing-1",
			in:   []int{8, 7, 6, 5, 4},
			exp:  2,
		},
		{
			name: "ever-increasing",
			in:   []int{5, 6, 7, 8},
			exp:  1,
		},
		{
			name: "ever-increasing-1",
			in:   []int{5, 6, 7, 8, 9},
			exp:  2,
		},
		{
			name: "ever-increasing-1",
			in:   []int{5, 5, 5, 5, 5, 5},
			exp:  2,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := minChopCount(tc.in)
			assert.Equal(t, tc.exp, out)
		})
	}
}

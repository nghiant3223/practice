package longestincreasingsubsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// [10,9,2,5,3,7,101,18]
// [0,1,0,3,2,3]
// [7,7,7,7,7,7,7]
// [-3, 10, 5, 12, 15]

func TestHappy(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   []int
		out  int
	}{
		{
			name: "1",
			in:   []int{10, 9, 2, 5, 3, 7, 101, 18},
			out:  4,
		},
		{
			name: "2",
			in:   []int{0, 1, 0, 3, 2, 3},
			out:  4,
		},
		{
			name: "3",
			in:   []int{7, 7, 7, 7, 7, 7, 7},
			out:  1,
		},
		{
			name: "4",
			in:   []int{-3, 10, 5, 12, 15},
			out:  4,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := lengthOfLIS(tc.in)
			assert.Equal(t, tc.out, out)
		})
	}
}

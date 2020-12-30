package houserobber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHappy(t *testing.T) {
	for _, tc := range []struct {
		in  []int
		out int
	}{
		{
			in:  []int{1},
			out: 1,
		},
		{
			in:  []int{1, 2, 3, 1},
			out: 4,
		},
		{
			in:  []int{2, 1, 1, 3},
			out: 4,
		},
		{
			in:  []int{2, 7, 9, 3, 1},
			out: 12,
		},
		{
			in:  []int{1, 2},
			out: 2,
		},
	} {
		t.Run("", func(t *testing.T) {
			out := rob(tc.in)
			assert.Equal(t, tc.out, out)
		})
	}
}

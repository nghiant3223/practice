package rmduplicatefromsortedarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHappy(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   []int
		out  int
	}{
		{
			name: "1",
			in:   []int{1, 2, 2},
			out:  2,
		},
		{
			name: "2",
			in:   []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			out:  5,
		},
	} {
		out := rmDuplicatesFromSortedArray(tc.in)
		assert.Equal(t, tc.out, out)
	}
}

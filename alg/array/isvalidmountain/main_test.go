package isvalidmountain

import (
	"testing"

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
			in:   []int{1, 3, 2},
			out:  true,
		},
		{
			name: "2",
			in:   []int{1, 2, 3},
			out:  false,
		},
		{
			name: "3",
			in:   []int{4, 3, 1, 5, 6},
			out:  false,
		},
		{
			name: "4",
			in:   []int{1, 5, 7, 2, 3, 4},
			out:  false,
		},
		{
			name: "5",
			in:   []int{9, 6, 5, 3, 2},
			out:  false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := isValidMountain(tc.in)
			assert.Equal(t, tc.out, out)
		})
	}
}

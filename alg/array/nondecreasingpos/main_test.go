package nondecreasingpos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHappy(t *testing.T) {
	type testcase struct {
		name string
		in   []int
		out  bool
	}
	for _, tc := range []testcase{
		{
			name: "1",
			in:   []int{1, 4, 2, 3},
			out:  true,
		},
		{
			name: "1",
			in:   []int{4, 2, 3},
			out:  true,
		},
		{
			name: "2",
			in:   []int{4, 2, 1},
			out:  false,
		},
		{
			name: "3",
			in:   []int{5, 4, 6, 7, 8, 9, 3},
			out:  false,
		},
		{
			name: "4",
			in:   []int{5, 4, 6, 7, 2, 9, 3},
			out:  false,
		},
		{
			name: "5",
			in:   []int{},
			out:  true,
		},
		{
			name: "6",
			in:   []int{1},
			out:  true,
		},
		{
			name: "7",
			in:   []int{1, 1, 1, 1, 1},
			out:  true,
		},
		{
			name: "8",
			in:   []int{1, 2, 3, 4, 5, 6, 7},
			out:  true,
		},
		{
			name: "9",
			in:   []int{3, 4, 2, 3},
			out:  false,
		},
		{
			name: "10",
			in:   []int{5, 7, 1, 8},
			out:  true,
		},
		{
			name: "11",
			in:   []int{5, 7, 8, 9},
			out:  true,
		},
		{
			name: "12",
			in:   []int{1, 3, 4, 2, 3},
			out:  false,
		},
		{
			name: "13",
			in:   []int{1, 3, 4, 2, 2},
			out:  false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ans := checkPossibility(tc.in)
			assert.Equal(t, tc.out, ans)
		})
	}
}

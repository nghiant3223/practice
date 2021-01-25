package allanagrams

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHappy(t *testing.T) {
	for _, tc := range []struct {
		name string
		in1  string
		in2  string
		out  []int
	}{
		{
			name: "1",
			in1:  "abacebca",
			in2:  "abc",
			out:  []int{1, 5},
		},
		{
			name: "2",
			in1:  "cbaebabacd",
			in2:  "abc",
			out:  []int{0, 6},
		},
	} {
		out := allAnagrams(tc.in1, tc.in2)
		assert.Equal(t, tc.out, out)
	}
}

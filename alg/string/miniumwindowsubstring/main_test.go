package miniumwindowsubstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHappy(t *testing.T) {
	for _, tc := range []struct {
		name string
		in1  string
		in2  string
		out  string
	}{
		{
			name: "1",
			in1:  "ADOBECODEBANC",
			in2:  "ABC",
			out:  "BANC",
		},
	} {
		out := minimumWindowSubstring(tc.in1, tc.in2)
		assert.Equal(t, tc.out, out)
	}
}

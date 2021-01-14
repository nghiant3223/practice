package cansumtarget

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHappy(t *testing.T) {
	for _, tc := range []struct {
		name string
		in1  int
		in2  []int
		out  bool
	}{
		{
			name: "1",
			in1:  7,
			in2:  []int{5, 3, 4, 1},
			out:  true,
		},
		{
			name: "2",
			in1:  7,
			in2:  []int{3, 1},
			out:  true,
		},
		{
			name: "3",
			in1:  7,
			in2:  []int{2, 4},
			out:  false,
		},
		{
			name: "4",
			in1:  7,
			in2:  []int{2, 4},
			out:  false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := canSum(tc.in1, tc.in2)
			assert.Equal(t, tc.out, out)
		})
	}
}

package bestsumtarget

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHappy(t *testing.T) {
	for _, tc := range []struct {
		name string
		in1  int
		in2  []int
		out  []int
	}{
		{
			name: "1",
			in1:  8,
			in2:  []int{2, 3, 7, 6},
			out:  []int{2, 6},
		},
		{
			name: "2",
			in1:  8,
			in2:  []int{2, 3, 7, 8, 6},
			out:  []int{8},
		},
		{
			name: "3",
			in1:  7,
			in2:  []int{5, 3, 4, 7},
			out:  []int{7},
		},
		{
			name: "4",
			in1:  8,
			in2:  []int{2, 3, 5},
			out:  []int{3, 5},
		},
		{
			name: "5",
			in1:  8,
			in2:  []int{1, 4, 5},
			out:  []int{4, 4},
		},
		{
			name: "6",
			in1:  100,
			in2:  []int{1, 2, 5, 25},
			out:  []int{25, 25, 25, 25},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := bestSum(tc.in1, tc.in2)
			for _, number := range out {
				assert.Contains(t, tc.out, number)
			}
		})
	}
}

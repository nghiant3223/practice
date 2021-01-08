package checkarrayformation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanFormArray(t *testing.T) {
	for _, tc := range []struct {
		name string
		in1  []int
		in2  [][]int
		out  bool
	}{
		{
			name: "1",
			in1:  []int{91, 4, 64, 78},
			in2:  [][]int{{78}, {4, 64}, {91}},
			out:  true,
		},
		{
			name: "2",
			in1: []int{49,18,16},
			in2: [][]int{{16,18,49}},
			out: false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := canFormArray(tc.in1, tc.in2)
			assert.Equal(t, tc.out, out)
		})
	}
}

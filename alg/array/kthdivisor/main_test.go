package kthdivisor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHappy(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   [2]int
		out  int
	}{
		{
			name: "1",
			in:   [2]int{12, 3},
			out:  3,
		},
		{
			name: "2",
			in:   [2]int{7, 2},
			out:  7,
		},
		{
			name: "3",
			in:   [2]int{4, 4},
			out:  -1,
		},
		{
			name: "4",
			in:   [2]int{1, 1},
			out:  1,
		},
		{
			name: "5",
			in:   [2]int{1000, 3},
			out:  4,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := kthFactor(tc.in[0], tc.in[1])
			assert.Equal(t, tc.out, out)
		})
	}
}

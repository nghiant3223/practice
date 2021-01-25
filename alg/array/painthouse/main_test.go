package painthouse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHappy(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   [][3]int
		out  int
	}{
		{
			name: "1",
			in: [][3]int{
				{17, 2, 17},
				{16, 16, 5},
				{14, 3, 19},
			},
			out: 10,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := paintHouse(tc.in)
			assert.Equal(t, tc.out, out)
		})
	}
}

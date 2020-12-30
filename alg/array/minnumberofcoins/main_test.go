package minnumberofcoins

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHappy(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   int
		out  int
	}{
		{
			name: "3",
			in:   3,
			out:  2,
		},
		{
			name: "6",
			in:   6,
			out:  2,
		},
		{
			name: "10",
			in:   10,
			out:  2,
		},
		{
			name: "12",
			in:   12,
			out:  3,
		},
		{
			name: "13",
			in:   13,
			out:  4,
		},
		{
			name: "28",
			in:   28,
			out:  7,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := minNumberOfCoins(tc.in)
			assert.Equal(t, tc.out, out)
		})
	}
}

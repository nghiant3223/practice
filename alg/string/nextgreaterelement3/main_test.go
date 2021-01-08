package nextgreaterelement3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTable(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   int
		out  int
	}{
		{
			name: "1",
			in:   218765,
			out:  251678,
		},
		{
			name: "2",
			in:   123456,
			out:  123465,
		},
		{
			name: "3",
			in:   4321,
			out:  -1,
		},
		{
			name: "4",
			in:   534976,
			out:  536479,
		},
		{
			name: "5",
			in:   11,
			out:  -1,
		},
		{
			name: "6",
			in:   230241,
			out:  230412,
		},
		{
			name: "7",
			in:   1753,
			out:  3157,
		},
		{
			name: "8",
			in:   5753,
			out:  7355,
		},
		{
			name: "9",
			in:   7753,
			out:  -1,
		},
		{
			name: "9",
			in:   17753,
			out:  31577,
		},
		{
			name: "10",
			in:   3486,
			out:  3648,
		},
		{
			name: "11",
			in:   12,
			out:  21,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := nextGreaterElement(tc.in)
			assert.Equal(t, tc.out, out)
		})
	}
}

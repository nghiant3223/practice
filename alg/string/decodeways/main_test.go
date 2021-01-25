package decodeways

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHappy(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   string
		out  int
	}{
		{
			name: "1",
			in:   "12345",
			out:  3,
		},
		{
			name: "2",
			in:   "12",
			out:  2,
		},
		{
			name: "3",
			in:   "226",
			out:  3,
		},
		{
			name: "4",
			in:   "0000",
			out:  0,
		},
		{
			name: "5",
			in:   "3020",
			out:  0,
		},
		{
			name: "6",
			in:   "2010",
			out:  1,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := numDecodings(tc.in)
			assert.Equal(t, tc.out, out)
		})
	}
}

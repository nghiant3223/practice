package sortcharbyfrequency

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHappy(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "1",
			in:   "tree",
			out:  "eert",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := frequencySort(tc.in)
			assert.Equal(t, tc.out, out)
		})
	}

}

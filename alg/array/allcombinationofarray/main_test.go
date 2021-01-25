package allcombinationofarray

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHappy(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   []int
		out  [][]int
	}{
		{
			name: "1",
			in:   []int{1, 2, 3},
		},
	} {
		out := allCombinationOfArray(tc.in)
		log.Println(out)
		assert.Equal(t, tc.out, out)
	}
}

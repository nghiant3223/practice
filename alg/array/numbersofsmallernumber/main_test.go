package numbersofsmallernumber

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHappy(t *testing.T) {
	type testcase struct {
		in  []int
		out []int
	}
	for i, tc := range []testcase{
		{
			in:  []int{8, 1, 2, 2, 3},
			out: []int{4, 0, 1, 1, 3},
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			ans := numberOfSmallerNumbers(tc.in)
			assert.Equal(t, tc.out, ans)
		})
	}
}

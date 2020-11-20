package alphabetwordlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHappy(t *testing.T) {
	type testcase struct {
		name string
		k, n int
		exp  []string
	}
	for _, tc := range []testcase{
		{
			name: "1",
			k:    3,
			n:    5,
			exp:  []string{"aba", "abc", "aca", "acb", "bab"},
		},
		{
			name: "1",
			k:    2,
			n:    1,
			exp:  []string{"ab"},
		},
		{
			name: "1",
			k:    2,
			n:    2,
			exp:  []string{"ab", "ac"},
		},
		{
			name: "1",
			k:    2,
			n:    3,
			exp:  []string{"ab", "ac", "ba"},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := alphabetWordList(tc.k, tc.n)
			assert.Equal(t, tc.exp, out)
			ans = []string{}
		})
	}
}

package canconstructstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHappy(t *testing.T) {
	for _, tc := range []struct {
		name string
		in1  string
		in2  []string
		out  bool
	}{
		{
			name: "1",
			in1:  "abcdef",
			in2:  []string{"ab", "abc", "cd", "def", "abcd"},
			out:  true,
		},
		{
			name: "2",
			in1:  "skateboard",
			in2:  []string{"bo", "rd", "ate", "t", "ska", "sh", "boar"},
			out:  false,
		},
		{
			name: "3",
			in1:  "yooaioyy",
			in2:  []string{"y", "i", "o", "ai", "a"},
			out:  true,
		},
		{
			name: "4",
			in1:  "enterapotentpot",
			in2:  []string{"a", "p", "ent", "enter", "ot", "o", "t"},
			out:  true,
		},
		{
			name: "5",
			in1:  "eeeee",
			in2:  []string{"eee", "eeee"},
			out:  false,
		},
		{
			name: "6",
			in1:  "eeeee",
			in2:  []string{"ee", "eee"},
			out:  true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := canConstruct(tc.in1, tc.in2)
			assert.Equal(t, tc.out, out)
		})
	}
}

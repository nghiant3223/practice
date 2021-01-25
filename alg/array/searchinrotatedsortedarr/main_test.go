package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCutIndex(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   []int
		out  int
	}{
		{
			name: "1",
			in:   []int{4, 5, 6, 7, 0, 1, 2},
			out:  7,
		},
		{
			name: "2",
			in:   []int{4, 5, 6, 7, 0, 1, 2, 3},
			out:  7,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := findCutIndex(tc.in)
			assert.Equal(t, tc.out, out)
		})
	}
}

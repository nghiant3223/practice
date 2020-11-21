package amounttotext

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHappy(t *testing.T) {
	type testcase struct {
		in  int
		out string
	}
	for i, tc := range []testcase{
		{
			in:  2,
			out: "hai",
		},
		{
			in:  12,
			out: "mười hai",
		},
		{
			in:  10,
			out: "mười",
		},
		{
			in:  22,
			out: "hai mươi hai",
		},
		{
			in:  20,
			out: "hai mươi",
		},
		{
			in:  19,
			out: "mười chín",
		},
		{
			in:  230,
			out: "hai trăm ba mươi",
		},
		{
			in:  236,
			out: "hai trăm ba mươi sáu",
		},
		{
			in:  200,
			out: "hai trăm",
		},
		{
			in:  299,
			out: "hai trăm chín mươi chín",
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			ans := amountToText(tc.in)
			assert.Equal(t, tc.out, ans)
		})
	}
}

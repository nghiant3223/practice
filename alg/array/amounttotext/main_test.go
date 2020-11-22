package amounttotext

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHappy(t *testing.T) {
	type testcase struct {
		in  string
		out string
	}
	for _, tc := range []testcase{
		{
			in:  "2",
			out: "hai",
		},
		{
			in:  "12",
			out: "mười hai",
		},
		{
			in:  "10",
			out: "mười",
		},
		{
			in:  "22",
			out: "hai mươi hai",
		},
		{
			in:  "20",
			out: "hai mươi",
		},
		{
			in:  "19",
			out: "mười chín",
		},
		{
			in:  "230",
			out: "hai trăm ba mươi",
		},
		{
			in:  "236",
			out: "hai trăm ba mươi sáu",
		},
		{
			in:  "200",
			out: "hai trăm",
		},
		{
			in:  "203",
			out: "hai trăm linh ba",
		},
		{
			in:  "299",
			out: "hai trăm chín mươi chín",
		},
		{
			in:  "1000",
			out: "một nghìn",
		},
		{
			in:  "1001",
			out: "một nghìn không trăm linh một",
		},
		{
			in:  "1030",
			out: "một nghìn không trăm ba mươi",
		},
		{
			in:  "1039",
			out: "một nghìn không trăm ba mươi chín",
		},
		{
			in:  "1010",
			out: "một nghìn không trăm mười",
		},
		{
			in:  "1013",
			out: "một nghìn không trăm mười ba",
		},
		{
			in:  "1099",
			out: "một nghìn không trăm chín mươi chín",
		},
		{
			in:  "1999",
			out: "một nghìn chín trăm chín mươi chín",
		},
		{
			in:  "0001999",
			out: "một nghìn chín trăm chín mươi chín",
		},
		{
			in:  "92092",
			out: "chín mươi hai nghìn không trăm chín mươi hai",
		},
		{
			in:  "192092",
			out: "một trăm chín mươi hai nghìn không trăm chín mươi hai",
		},
		{
			in:  "1192092",
			out: "một triệu một trăm chín mươi hai nghìn không trăm chín mươi hai",
		},
		{
			in:  "1000000",
			out: "một triệu",
		},
		{
			in:  "1000001",
			out: "một triệu không trăm linh một",
		},
		{
			in:  "1001001",
			out: "một triệu một nghìn không trăm linh một",
		},
	} {
		t.Run(fmt.Sprint(tc.in), func(t *testing.T) {
			ans := AmountToText(tc.in)
			assert.Equal(t, tc.out, ans)
		})
	}
}

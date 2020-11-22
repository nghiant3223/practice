package amounttotext

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	zeroChar = '0'
)

const (
	zero   = "0"
	one    = "1"
	two    = "2"
	three  = "3"
	four   = "4"
	five   = "5"
	six    = "6"
	seven  = "7"
	eight  = "8"
	nine   = "9"
	ten    = "10"
	twenty = "20"
)

const (
	tenSuffix      = "0"
	hundredSuffix  = "00"
	thousandSuffix = "000"
	millionSuffix  = "000000"
)

var digitToText = map[string]string{
	zero:  "không",
	one:   "một",
	two:   "hai",
	three: "ba",
	four:  "bốn",
	five:  "năm",
	six:   "sáu",
	seven: "bảy",
	eight: "tám",
	nine:  "chín",
}

func AmountToText(amount string) string {
	if !validateAmount(amount) {
		return ""
	}
	if isZeroSeries(amount) {
		return zero
	}
	amount = strings.TrimLeft(amount, zero)
	return amountToText(amount)
}

func validateAmount(amount string) bool {
	if amount == "" {
		return false
	}
	_, err := strconv.Atoi(amount)
	if err != nil {
		return false
	}
	return true
}

func isZeroSeries(amount string) bool {
	for _, c := range amount {
		if c != zeroChar {
			return false
		}
	}
	return true
}

func amountToText(amount string) string {
	length := len(amount)
	if length == 0 {
		return ""
	}
	if length == 1 {
		return digitToText[amount]
	}
	if length == 2 {
		if amount == ten {
			return "mười"
		}
		if amount < twenty {
			return fmt.Sprintf("mười %s", AmountToText(amount[1:]))
		}
		if strings.HasSuffix(amount, tenSuffix) {
			return fmt.Sprintf("%s mươi", digitToText[string(amount[0])])
		}
		return fmt.Sprintf("%s mươi %s", digitToText[string(amount[0])], amountToText(amount[1:]))
	}
	if length == 3 {
		if strings.HasSuffix(amount, hundredSuffix) {
			return fmt.Sprintf("%s trăm", digitToText[string(amount[0])])
		}
		if string(amount[1]) == "0" {
			return fmt.Sprintf("%s trăm linh %s", digitToText[string(amount[0])], digitToText[string(amount[2])])
		}
		return fmt.Sprintf("%s trăm %s", digitToText[string(amount[0])], amountToText(amount[1:]))
	}
	if length <= 6 {
		if strings.HasSuffix(amount, thousandSuffix) {
			return fmt.Sprintf("%s nghìn", digitToText[string(amount[0])])
		}
		return fmt.Sprintf("%s nghìn %s", amountToText(amount[:length-3]), amountToText(amount[length-3:]))
	}
	if length <= 9 {
		if strings.HasSuffix(amount, millionSuffix) {
			return fmt.Sprintf("%s triệu", amountToText(amount[:length-6]))
		}
		if amount[length-6:length-3] == thousandSuffix {
			return fmt.Sprintf("%s triệu %s", amountToText(amount[:length-6]), amountToText(amount[length-3:]))
		}
		return fmt.Sprintf("%s triệu %s", amountToText(amount[:length-6]), amountToText(amount[length-6:]))
	}
	return ""
}

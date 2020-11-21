package amounttotext

import (
	"fmt"
	"math"
)

var mapDigitToText = map[int]string{
	0: "không",
	1: "một",
	2: "hai",
	3: "ba",
	4: "bốn",
	5: "năm",
	6: "sáu",
	7: "bảy",
	8: "tám",
	9: "chín",
}

func amountToText(amount int) string {
	if amount < 10 {
		return mapDigitToText[amount]
	}
	if amount == 10 {
		return "mười"
	}
	if amount < 20 {
		return fmt.Sprintf("mười %s", amountToText(reduce(amount)))
	}
	if amount < 100 {
		if amount%10 == 0 {
			return fmt.Sprintf("%s mươi", mapDigitToText[amount/10])
		}
		return fmt.Sprintf("%s mươi %s", mapDigitToText[amount/10], amountToText(reduce(amount)))
	}
	if amount < 1000 {
		if amount%100 == 0 {
			return fmt.Sprintf("%s trăm", mapDigitToText[amount/100])
		}
		return fmt.Sprintf("%s trăm %s", mapDigitToText[amount/100], amountToText(reduce(amount)))
	}
	if amount < 100000 {
		return fmt.Sprintf("%s nghìn %s", )

	}
	return ""
}

func reduce(number int) int {
	return number % int(math.Pow10(int(math.Log10(float64(number)))))
}

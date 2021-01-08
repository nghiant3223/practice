package nextgreaterelement3

import "strconv"

func nextGreaterElement(n int) int {
	number := strconv.Itoa(n)
	numberLength := len(number)

	pivot := -1
	lastIndex := len(number) - 1
	decreasingPart := string(number[lastIndex])
	for i := lastIndex - 1; i >= 0; i-- {
		if number[i] < decreasingPart[0] {
			pivot = i
			break
		}
		decreasingPart = string(number[i]) + decreasingPart
	}

	decreasingPartLength := len(decreasingPart)
	if decreasingPartLength == numberLength {
		return -1
	}

	unchangedPart := number[:pivot]
	stringAnswer := unchangedPart

	pivotElement := number[pivot]
	candidate, index := findGreaterThanPivot(decreasingPart, pivotElement)
	decreasingPartBytes := []byte(decreasingPart)
	decreasingPartBytes[index] = pivotElement
	stringAnswer += string(candidate)
	stringAnswer = concatReverse(stringAnswer, string(decreasingPartBytes))

	answer, err := strconv.Atoi(stringAnswer)
	if err != nil {
		return 0
	}
	if answer > 2147483647 {
		return -1
	}
	return answer
}

func findGreaterThanPivot(decreasingPart string, pivotElement byte) (byte, int) {
	for i := len(decreasingPart) - 1; i >= 0; i-- {
		if decreasingPart[i] > pivotElement {
			return decreasingPart[i], i
		}
	}
	return 0, -1
}

func concatReverse(a, b string) string {
	for i := len(b) - 1; i >= 0; i-- {
		a += string(b[i])
	}
	return a
}

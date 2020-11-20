package numbersofsmallernumber

const (
	lowerBound = 0
	upperBound = 100
	length     = upperBound - lowerBound + 1
)

func numberOfSmallerNumbers(numbers []int) []int {
	countings := make([]int, length)
	for _, number := range numbers {
		countings[number]++
	}
	accumulative := 0
	for i := range countings {
		tmp := countings[i]
		countings[i] = accumulative
		accumulative += tmp
	}
	for i, n := range numbers {
		numbers[i] = countings[n]
	}
	return numbers
}

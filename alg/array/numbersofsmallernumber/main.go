package numbersofsmallernumber

import (
	"sort"
)

func numberOfSmallerNumbers(numbers []int) []int {
	m := buildFrequencyMap(numbers)
	s := buildSliceSortByKey(m)
	accumulativeCount := 0
	for _, number := range s {
		count := m[number]
		m[number] = accumulativeCount
		accumulativeCount += count
	}
	for i, n := range numbers {
		numbers[i] = m[n]
	}
	return numbers
}

func buildFrequencyMap(numbers []int) map[int]int {
	m := make(map[int]int, len(numbers))
	for _, n := range numbers {
		m[n]++
	}
	return m
}

func buildSliceSortByKey(m map[int]int) []int {
	s := make([]int, len(m))
	i := 0
	for k := range m {
		s[i] = k
		i++
	}
	sort.Ints(s)
	return s
}

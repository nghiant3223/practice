package main

import (
	"log"
)

type Element struct {
	Value     int
	Frequency int
}

func topKFrequent(nums []int, k int) []int {
	mapNumToElement := buildFrequencyMap(nums)
	elements := make([]Element, len(mapNumToElement))
	i := 0
	for _, element := range mapNumToElement {
		elements[i] = element
		i++
	}
	minHeap, restElements := buildHeapWithSizeK(elements, k)
	iterationCount := len(restElements)
	for i := 0; i < iterationCount; i++ {
		minFrequency := minHeap[0].Frequency
		if restElements[i].Frequency <= minFrequency {
			continue
		}
		minHeap[0] = restElements[i]
		reheapDown(minHeap, 0)
	}
	ans := make([]int, k)
	for i := 0; i < k; i++ {
		ans[i] = minHeap[i].Value
	}
	return ans
}

func buildFrequencyMap(nums []int) map[int]Element {
	mapNumToElement := make(map[int]Element)
	for _, num := range nums {
		element, ok := mapNumToElement[num]
		if !ok {
			element = Element{
				Value:     num,
				Frequency: 0,
			}
		}
		element.Frequency++
		mapNumToElement[num] = element
	}
	return mapNumToElement
}

func buildHeapWithSizeK(elements []Element, k int) (heap []Element, rest []Element) {
	for i := 0; i < k; i++ {
		reheapUp(elements, i)
	}
	return elements[:k], elements[k:]
}

func reheapUp(elements []Element, i int) {
	if i < 0 {
		return
	}
	parentIndex := (i - 1) / 2
	if elements[i].Frequency < elements[parentIndex].Frequency {
		swap(elements, i, parentIndex)
		reheapUp(elements, parentIndex)
	}
}

func reheapDown(elements []Element, i int) {
	length := len(elements)
	leftIndex := 2*i + 1
	rightIndex := 2*i + 2
	if leftIndex >= length {
		return
	}
	var maxIndex int
	if leftIndex == length-1 {
		maxIndex = leftIndex
	} else {
		left := elements[leftIndex]
		right := elements[rightIndex]
		if left.Frequency < right.Frequency {
			maxIndex = leftIndex
		} else {
			maxIndex = rightIndex
		}
	}
	if elements[maxIndex].Frequency < elements[i].Frequency {
		swap(elements, i, maxIndex)
		reheapDown(elements, maxIndex)
	}
}

func swap(elements []Element, i, j int) {
	elements[i], elements[j] = elements[j], elements[i]
}

func main() {
	nums := []int{5, -3, 9, 1, 7, 7, 9, 10, 2, 2, 10, 10, 3, -1, 3, 7, -9, -1, 3, 3}
	k := 3
	ans := topKFrequent(nums, k)
	log.Print(ans)
}

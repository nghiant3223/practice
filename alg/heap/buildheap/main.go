package main

import "log"

func buildHeap(nums []int) {
	l := len(nums)
	iterationCount := l / 2
	for i := 0; i < iterationCount; i++ {
		swap(nums, i, l-1-i)
		reheapDown(nums, i)
	}
}

func reheapDown(nums []int, i int) {
	length := len(nums)
	leftIndex := 2*i + 1
	rightIndex := 2*i + 2
	if leftIndex >= length {
		return
	}
	if leftIndex == length {
		if nums[leftIndex] < nums[i] {
			swap(nums, i, leftIndex)
			return
		}
		return
	}
	left := nums[leftIndex]
	right := nums[rightIndex]
	var minIndex int
	if left < right {
		minIndex = leftIndex
	} else {
		minIndex = rightIndex
	}
	if nums[minIndex] < nums[i] {
		swap(nums, i, minIndex)
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func main() {
	nums := []int{4, 10, 3, 5, 1}
	buildHeap(nums)
	log.Print(nums)
}

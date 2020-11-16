package main

import "fmt"

// You are given an integer array nums sorted in ascending order, and an integer target.
// Suppose that nums is rotated at some pivot unknown to you beforehand (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
// If target is found in the array return its index, otherwise, return -1.

func searchInRotatedSortedArray(nums []int, target int) int {
	v := findValley(nums)
	findLeft := binarySearch(nums, target, 0, v-1)
	if findLeft != -1 {
		return findLeft
	}
	findRight := binarySearch(nums, target, v, len(nums)-1)
	return findRight
}

func binarySearch(nums []int, target int, left, right int) int {
	for left <= right {
		m := (left + right) / 2
		if nums[m] == target {
			return m
		}
		if nums[m] < target {
			left = m + 1
		} else {
			right = m - 1
		}
	}
	return - 1
}

func findValley(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left != right {
		if left+1 == right {
			if nums[left] > nums[right] {
				return right
			}
			return -1
		}
		mid := (left + right) / 2
		if nums[left] > nums[mid] {
			right = mid
		} else if nums[mid] > nums[right] {
			left = mid
		} else {
			return -1
		}
	}
	return left
}

// 1 2 3 4 0
//
//0 1 2 3 4
//
//3 4 0 1 2
//
//4 0 1 2 3
//
//2 3 4 0 1

func main() {
	nums := []int{2, 3, 4, 0, 1}
	fmt.Println(searchInRotatedSortedArray(nums, 1))
}

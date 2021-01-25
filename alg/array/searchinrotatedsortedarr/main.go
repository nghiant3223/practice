package main

// You are given an integer array nums sorted in ascending order, and an integer target.
// Suppose that nums is rotated at some pivot unknown to you beforehand (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
// If target is found in the array return its index, otherwise, return -1.

func searchInRotatedSortedArray(nums []int, target int) int {
	v := findCutIndex(nums)
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
	return -1
}

func findCutIndex(nums []int) int {
	if len(nums) < 3 {
		return -1
	}
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		if nums[m-1] < nums[m] && nums[m] > nums[m+1] {
			return nums[m]
		}
		if nums[m] > nums[l] {
			l = m
		} else {
			r = m
		}
	}
	return -1
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

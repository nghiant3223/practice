package main

import (
	"fmt"
)

func canJump(nums []int) bool {
	return recur(nums, len(nums), 0)
}

func recur(nums []int, length int, i int) bool {
	if i+nums[i] >= length-1 {
		return true
	}
	for j := 1; j <= nums[i]; j++ {
		if recur(nums, length, i+j) {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{3, 2, 1, 0, 4}
	ans := canJump(nums)
	fmt.Print(ans)
}

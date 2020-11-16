package main

import "fmt"

// Given an array nums of n integers where n > 1, return an array output such that
// output[i] is equal to the product of all the elements of nums except nums[i].

func productOfArrayExceptSelf(nums []int) []int {
	cnt := len(nums)
	forward := make([]int, cnt)
	pForward := 1
	forward[0] = pForward
	for i := 1; i < cnt; i++ {
		forward[i] = pForward
		pForward = pForward * nums[i]
	}
	backward := make([]int, cnt)
	pBackward := 1
	backward[0] = pBackward
	for i := cnt - 1; i >= 0; i-- {
		backward[i] = pBackward
		pBackward = pBackward * nums[i]
	}
	ans := make([]int, cnt)
	for i := 0; i < len(nums); i++ {
		ans[i] = forward[i] * backward[i]
	}
	return ans
}

func main() {
	ans := productOfArrayExceptSelf([]int{1, 2, 3, 4})
	fmt.Println(ans)
}

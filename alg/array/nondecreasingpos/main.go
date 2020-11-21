package nondecreasingpos

func checkPossibility(nums []int) bool {
	length := len(nums)
	if length < 2 {
		return true
	}
	count := 0
	index := -1
	for i := 0; i < length-1; i++ {
		if nums[i] > nums[i+1] {
			count++
			if count >= 2 {
				return false
			}
			index = i
		}
	}
	if count == 0 {
		return true
	}
	if count == 1 {
		// When encounter a decreasing pair of numbers, such as: a b c 4 2 x y z
		// There a two possibilities to make array non-decreasing, removing 4 or removing 2
		// To remove 4, we need to ensure that c <= 2 to make it possible, eg: 1 4 2 3
		// To remove 2, we need to ensure that 4 <= x to make it possible, eg: 3 4 2 8
		// If one of the possibilities is possible, return true
		if index+2 >= length || nums[index] <= nums[index+2] {
			return true
		}
		if index-1 < 0 || nums[index-1] <= nums[index+1] {
			return true
		}
	}
	return false
}

package isvalidmountain

func isValidMountain(nums []int) bool {
	count := len(nums)
	if count < 3 {
		return false
	}
	isIncreasing := nums[1] > nums[0]
	if !isIncreasing {
		return false
	}
	isDecreasing := false
	for i := 2; i < count; i++ {
		if isDecreasing && nums[i-1] < nums[i] {
			return false
		}
		if nums[i-1] > nums[i] {
			isDecreasing = true
		}
	}
	return isIncreasing && isDecreasing
}

package rmduplicatefromsortedarray

// Given a sorted array nums, remove duplicates in-place such that each element appear only once and return the new length
// Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory

func rmDuplicatesFromSortedArray(nums []int) int {
	count := len(nums)
	if count < 2 {
		return count
	}
	prev, cur := 0, 1
	for cur < count {
		for cur < count && nums[cur] == nums[prev] {
			cur++
		}
		if cur >= count {
			break
		}
		nums[prev+1] = nums[cur]
		prev++
		cur++
	}
	return prev + 1
}

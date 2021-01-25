package topklargest

func reheapDown(nums []int, parent int) {
	count := len(nums)
	candidate := parent
	left := 2*parent + 1
	right := 2*parent + 2
	if left < count && nums[left] > nums[candidate] {
		candidate = left
	}
	if right < count && nums[right] > nums[candidate] {
		candidate = right
	}
	if candidate != parent {
		nums[candidate], nums[parent] = nums[parent], nums[candidate]
		reheapDown(nums, candidate)
	}
}

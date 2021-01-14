package ds

func ParseLinkedList(nums []int) *Node {
	count := len(nums)
	if count == 0 {
		return nil
	}
	head := NewNode(nums[0])
	tmp := head
	for i := 1; i < count; i++ {
		node := NewNode(nums[i])
		tmp.Next = node
		tmp = tmp.Next
	}
	return head
}

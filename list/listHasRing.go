package list

// 判断链表是否有环,逻辑是使用双指针，一快一慢，有环相遇，无环不相遇
func ListHasRing(head *ListNode) bool {

	// 判空
	if head == nil || head.Next == nil {
		return false
	}

	slowPoint := head
	fastPoint := head.Next

	for slowPoint != nil && fastPoint != nil {
		if slowPoint == fastPoint {
			return true
		}
		slowPoint = slowPoint.Next
		fastPoint = fastPoint.Next
		if fastPoint == nil {
			return false
		}
		fastPoint = fastPoint.Next
	}

	return false
}
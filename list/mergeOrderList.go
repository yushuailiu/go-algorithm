package list

// 合并有序链表，合并后仍然有序
func MergeOrderList(head1, head2 *ListNode) *ListNode {

	headNew := new(ListNode)
	cur := headNew

	for head1 != nil && head2 != nil  {
		for head1 != nil && head1.Val <= head2.Val  {
			cur.Next = head1
			cur = head1
			head1 = head1.Next
		}
		if head1 == nil {
			break
		}

		for head2 != nil && head2.Val <= head1.Val {
			cur.Next = head2
			cur = head2
			head2 = head2.Next
		}
	}

	if head1 != nil {
		cur.Next = head1
	}
	if head2 != nil {
		cur.Next = head2
	}

	return headNew.Next
}
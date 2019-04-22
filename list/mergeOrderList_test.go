package list

import (
	"testing"
)

var mergeOrderListList = [][][]int{
	{
		{1, 2, 3, 4},
		{7,8,9},
		{1, 2, 3, 4, 7, 8, 9},
	},{
		{1, 2, 3, 4},
		{3,8},
		{1, 2, 3, 3, 4, 8},
	},{
		{1, 2, 3, 4},
		{5},
		{1, 2, 3, 4, 5},
	},{
		{1, 2, 3, 4},
		{1},
		{1, 1, 2, 3, 4},
	},{
		{1, 3, 4},
		{2},
		{1, 2, 3, 4},
	},{
		{1, 3, 4},
		{},
		{1, 3, 4},
	},{
		{4, 10, 20},
		{1, 3},
		{1, 3, 4, 10, 20},
	},{
		{4, 10, 20},
		{5, 11, 24},
		{4, 5, 10, 11, 20, 24},
	},{
		{1, 3, 4},
		{0, 2, 4},
		{0, 1, 2, 3, 4, 4},
	},{
		{},
		{},
		{},
	},
}


func TestMergeOrderList(t *testing.T) {

	for _, items := range mergeOrderListList {
		head1 := generateList(items[0])
		head2 := generateList(items[1])

		headNew := MergeOrderList(head1, head2)

		for i := 0; i < len(items[2]); i++ {
			if headNew.Val != items[2][i] {
				t.Errorf("merger order list error %d > %d", headNew.Val, headNew.Next.Val)
			}
			headNew = headNew.Next
		}


		for headNew != nil && headNew.Next != nil {
			if headNew.Val > headNew.Next.Val {
				t.Errorf("merger order list error %d > %d", headNew.Val, headNew.Next.Val)
			}
			headNew = headNew.Next
		}

	}

}

func generateList(items []int) *ListNode {
	var head *ListNode
	cur := head
	for _, i := range items {
		temp := &ListNode{
			Val: i,
		}

		if head == nil {
			head = temp
			cur = temp
		} else {
			cur.Next = temp
			cur = temp
		}
	}

	return head
}
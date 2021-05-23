package list

import (
	"testing"
)

// 测试无环单向链表
func TestHasNotRingList(t *testing.T) {
	for i := 0; i < 100; i++ {
		head := generateNotRingList(i)
		hasRing := ListHasRing(head)
		if hasRing {
			t.Errorf("TestHasNotRingList error list length %d", i)
		}
	}
}

func TestHasRingList(t *testing.T) {
	for i := 1; i < 30; i++ {
		for j := 1; j <=i; j++ {
			head := generateRingList(i, j)
			hasRing := ListHasRing(head)
			if !hasRing {
				println("TestHasRingList length %d ring %d but not found ring %t", i, j, hasRing)
			}
		}
	}
}

// 生成指定长度无环链表
func generateNotRingList(total int) *ListNode {
	i := 0
	var head *ListNode
	var tail *ListNode
	for i < total {
		temp1 := &ListNode{
			Val:  i,
			Next: nil,
		}
		if head == nil {
			head = temp1
			tail = temp1
		} else {
			tail.Next = temp1
			tail = temp1
		}
		i++
	}
	return head
}

// 生成指定长度和指定环大小的有环链表
func generateRingList(total, ringNumber int) *ListNode {
	if total <= 0 || ringNumber <= 0 {
		panic("generateRingList total ringNumber can not less then zero")
	}
	i := 0
	var head *ListNode
	var tail *ListNode
	var kNote *ListNode
	var temp1 *ListNode
	for i < total {
		temp1 = &ListNode{
			Val:  i,
			Next: nil,
		}
		if head == nil {
			head = temp1
			tail = temp1
		} else {
			tail.Next = temp1
			tail = temp1
		}
		if i == total - ringNumber {
			kNote = temp1
		}
		i++
	}

	tail.Next = kNote

	return head
}
package list

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 该函数是实现从尾到头打印链表 —— 剑指offer 面试题六
// 一般从头到尾打印链表，直接遍历就行。这里若反向打印一般思路是先反转，或者遍历入堆，再pop堆打印即可。
// 这里实现使用递归的方式，打印本节点前先去检测子节点，正好实现反向打印节点。
func PrintListNodeReverse(header *ListNode) {
	if header.Next != nil {
		PrintListNodeReverse(header.Next)
	}
	fmt.Println(header.Val)
}

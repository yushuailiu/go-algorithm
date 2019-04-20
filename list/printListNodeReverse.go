package list

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type ReversePrintListNode func(node *ListNode)

// 该函数是实现从尾到头打印链表 —— 剑指offer 面试题六
// 一般从头到尾打印链表，直接遍历就行。这里若反向打印一般思路是先反转，或者遍历入堆，再pop堆打印即可。

// 递归打印
func ReversePrintListNodeRecursion(header *ListNode) {
	if header == nil {
		return
	}

	ReversePrintListNodeRecursion(header.Next)
	fmt.Println(header.Val)
}

// 循环遍历打印
func ReversePrintListNodeTraverse(header *ListNode) {
	var last *ListNode
	for cur := header; last != header; cur = header {
		for cur.Next != last {
			cur = cur.Next
		}
		fmt.Println(cur.Val)
		last = cur
	}
}

// 先反转再打印
func ReversePrintListNodeReverse(header *ListNode) {
	var pre *ListNode
	if header == nil {
		return
	}
	next := header.Next
	for next != nil {
		header.Next = pre
		pre = header
		header = next
		next = next.Next
	}
	header.Next = pre

	for header != nil {
		fmt.Println(header.Val)
		header = header.Next
	}
}
// 递归反转打印
func ReversePrintListNodeRecursionReverse(header *ListNode) {
	header = recursionReverse(nil, header)

	for header != nil {
		fmt.Println(header.Val)
		header = header.Next
	}
}

func recursionReverse(pre, header *ListNode) *ListNode {
	if header == nil {
		return pre
	}
	next := header.Next
	header.Next = pre

	return recursionReverse(header, next)
}

// 堆栈打印
func ReversePrintListNodeHeap(header *ListNode)  {
	heap := new(myHeap)
	for header != nil {
		heap.Push(header)
		header = header.Next
	}

	for heap.Len() > 0 {
		cur := heap.Pop()
		node := cur.(*ListNode)
		fmt.Println(node.Val)
	}
}

type myHeap []*ListNode

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i].Val < (*h)[j].Val
}

func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *myHeap) Len() int {
	return len(*h)
}

func (h *myHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *myHeap) Push(v interface{}) {
	*h = append(*h, v.(*ListNode))
}


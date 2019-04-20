package list

import (
	"fmt"
	"testing"
)

var list = [][]int{
	{1, 2, 3, 4},
	{1},
	{},
	{1, 2},
}

var ReversePrintListNodeList = map[string]ReversePrintListNode{
	"recursion":         ReversePrintListNodeRecursion,
	"traverse":          ReversePrintListNodeTraverse,
	"reverse":           ReversePrintListNodeReverse,
	"recursion reverse": ReversePrintListNodeRecursionReverse,
	"heap": ReversePrintListNodeHeap,
}

func TestPrintListNodeReverse(t *testing.T) {

	for _, items := range list {

		fmt.Println(items)

		for method, function := range ReversePrintListNodeList {
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
			fmt.Println(method + " start")
			function(head)
			fmt.Println(method + " end")

		}

	}

}

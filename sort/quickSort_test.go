package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {

	list := [10]int{22, 10, 1, 99, 22, 10, 11, 10, 2, 100}
	QuickSort(list[:])
	res := fmt.Sprintf("%v", list)

	if res != "[1 2 10 10 10 11 22 22 99 100]" {
		t.Error("quick sort", list, "should be [1 2 10 10 10 11 22 22 99 100] bu get",
			res)
	}

	list1 := [10]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	QuickSort(list1[:])
	res = fmt.Sprintf("%v", list1)
	if res != "[1 2 3 4 5 6 7 8 9 10]" {
		t.Error("quick sort", list, "should be [1 2 3 4 5 6 7 8 9 10] but get", list1)
	}

	list2 := [2]int{3, 2}
	QuickSort(list2[:])
	if list2[0] != 2 || list2[1] != 3 {
		t.Error("quick sort", list2, "fail")
	}

	list3 := [6]int{2, 2, 2, 2, 3, 2}
	QuickSort(list3[:])
	if list3[0] != 2 || list3[5] != 3 {
		t.Error("quick sort", list3, "fail")
	}
}

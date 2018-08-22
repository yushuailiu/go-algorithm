package search

import (
	"testing"
)

func TestBinary1(t *testing.T)  {

	list := [10]int{1, 20, 25, 30, 40, 50, 55, 67, 89, 102}

	if index := BinarySearch(list[:], 20); index != 1 {
		t.Error("search 20 in", list, "should be 1 but get", index)
	}

	if index := BinarySearch(list[:], 1); index != 0 {
		t.Error("search 1 in", list, "should be 0 but get", index)
	}

	if index := BinarySearch(list[:], 102); index != 9 {
		t.Error("search 102 in", list, "should be 9 but get", index)
	}

	if index := BinarySearch(list[:], 35); index != -1 {
		t.Error("search 35 in", list, "should be -1 but get", index)
	}

	if index := BinarySearch(list[:], -10); index != -1 {
		t.Error("search -10 in", list, "should be -1 but get", index)
	}
}

func TestBinary2(t *testing.T)  {

	list := [0]int{}

	if index := BinarySearch(list[:], 20); index != -1 {
		t.Error("search 20 in", list, "should be -1 but get", index)
	}

	list1 := [1]int{1}

	if index := BinarySearch(list1[:], 1); index != 0 {
		t.Error("search 1 in", list1, "should be 0 but get", index)
	}

	if index := BinarySearch(list1[:], 2); index != -1 {
		t.Error("search 2 in", list1, "should be -1 but get", index)
	}
}
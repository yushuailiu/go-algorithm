package sort

import (
	"testing"
)

var list = []map[string][]int{
	{
		"origin": {10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		"result": {1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
	{
		"origin": {22, 33, 22, 1, 2, 11, 22, 123, 110},
		"result": {1, 2, 11, 22, 22, 22, 33, 110, 123},
	},
	{
		"origin": {22, 23, 10, 9, 54, 17, 100, 2},
		"result": {2, 9, 10, 17, 22, 23, 54, 100},
	},
	{
		"origin": {22, 23, 22, 23, 22, 23, 22, 23},
		"result": {22, 22, 22, 22, 23, 23, 23, 23},
	},
	{
		"origin": {1, 1111, 298989, 8422, 2222, 44422},
		"result": {1, 1111, 2222, 8422, 44422, 298989},
	},
	{
		"origin": {22, 10, 1, 99, 22, 10, 11, 10, 2, 100},
		"result": {1, 2, 10, 10, 10, 11, 22, 22, 99, 100},
	},
	{
		"origin": {23, 22},
		"result": {22, 23},
	},
	{
		"origin": {1},
		"result": {1},
	},
	{
		"origin": {},
		"result": {},
	},
}

func listEqual(list1, list2 []int) bool {
	for index, _ := range list1 {
		if list1[index] != list2[index] {
			return false
		}
	}
	return true
}

func TestQuickSort2(t *testing.T) {
	for _, item := range list {
		origin := make([]int, len(item["origin"]))
		copy(origin, item["origin"])
		QuickSort(origin)
		if !listEqual(origin, item["result"]) {
			t.Error("quick sort", item["origin"], "should be", item["result"],
				"but get", origin)
		}
	}
}

func TestBubbleSort(t *testing.T) {
	for _, item := range list {
		origin := make([]int, len(item["origin"]))
		copy(origin, item["origin"])
		BubbleSort(origin)
		if !listEqual(origin, item["result"]) {
			t.Error("bubble sort", item["origin"], "should be", item["result"],
				"but get", origin)
		}
	}
}

func TestInsertionSort(t *testing.T) {
	for _, item := range list {
		origin := make([]int, len(item["origin"]))
		copy(origin, item["origin"])
		InsertionSort(origin)
		if !listEqual(origin, item["result"]) {
			t.Error("insertion sort", item["origin"], "should be", item["result"],
				"but get", origin)
		}
	}
}

func TestSelectionSort(t *testing.T) {
	for _, item := range list {
		origin := make([]int, len(item["origin"]))
		copy(origin, item["origin"])
		InsertionSort(origin)
		if !listEqual(origin, item["result"]) {
			t.Error("selection sort", item["origin"], "should be", item["result"],
				"but get", origin)
		}
	}
}
func TestCountSort(t *testing.T) {
	for _, item := range list {
		origin := make([]int, len(item["origin"]))
		copy(origin, item["origin"])
		InsertionSort(origin)
		if !listEqual(origin, item["result"]) {
			t.Error("count sort", item["origin"], "should be", item["result"],
				"but get", origin)
		}
	}
}

package sort

import (
	"testing"
)

func BenchmarkBubbleSort1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, item := range list {
			origin := make([]int, len(item["origin"]))
			copy(origin, item["origin"])
			BubbleSort1(origin)
			if !listEqual(origin, item["result"]) {
				b.Error("bubble sort", item["origin"], "should be", item["result"],
					"but get", origin)
			}
		}
	}
}
func BenchmarkBubbleSort2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, item := range list {
			origin := make([]int, len(item["origin"]))
			copy(origin, item["origin"])
			BubbleSort2(origin)
			if !listEqual(origin, item["result"]) {
				b.Error("bubble sort", item["origin"], "should be", item["result"],
					"but get", origin)
			}
		}
	}
}
func BenchmarkBubbleSort3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, item := range list {
			origin := make([]int, len(item["origin"]))
			copy(origin, item["origin"])
			BubbleSort3(origin)
			if !listEqual(origin, item["result"]) {
				b.Error("bubble sort", item["origin"], "should be", item["result"],
					"but get", origin)
			}
		}
	}
}

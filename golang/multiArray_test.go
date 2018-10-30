package golang

import (
	"math/rand"
	"testing"
)

var items = make([][]int32, 1000)

func init() {
	for i := 0; i < 1000; i++ {
		items[i] = make([]int32, 1000)
		for j := 0; j < 1000; j++ {
			items[i][j] = rand.Int31n(2)
		}
	}
}

// 横向遍历
func sumRows() int {
	var sum = 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			sum += int(items[i][j])
		}
	}
	return sum
}

// 纵向遍历
func sumCols() int {
	var sum = 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			sum += int(items[j][i])
		}
	}
	return sum
}

func BenchmarkRows(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumRows()
	}
}

func BenchmarkCols(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumCols()
	}
}

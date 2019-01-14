package golang

import "testing"

func nestFor1() {
	a := 0
	for i := 0; i < 10000000; i++ {
		for j := 0; j < 2; j++ {
			a = a + i
			a = a + j
		}
	}
}

func nestFor2() {
	a := 0
	for i := 0; i < 2; i++ {
		for j := 0; j < 10000000; j++ {
			a = a + i
			a = a + j
		}
	}
}

func BenchmarkNestFor1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nestFor1()
	}
}
func BenchmarkNestFor2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nestFor2()
	}
}

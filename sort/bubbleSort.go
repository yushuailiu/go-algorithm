package sort

func BubbleSort1(list []int) {
	n := len(list)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
}

func BubbleSort2(list []int) {
	n := len(list)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
}

func BubbleSort3(list []int) {
	n := len(list)
	flag := true

	for flag {
		flag = false
		for j := 0; j < n-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				flag = true
			}
		}
		n = n - 1
	}
}

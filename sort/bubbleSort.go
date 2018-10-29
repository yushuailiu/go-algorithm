package sort

// BubbleSort1 是冒泡排序的最近本实现
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

// BubbleSort2 是冒泡排序的一个简单地优化，每次可以少比较一次
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

// BubbleSort3 是对冒泡排序的一个终极优化，最极端的情况下只需要遍历一次
// 即可完成排序
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

package sort

// QuickSort implement quick sort
func QuickSort(list []int) {
	if len(list) <= 1 {
		return
	}

	left := 0
	right := len(list) - 1
	mid := list[0]

	for left < right {
		for list[right] > mid && left < right {
			right--
		}

		if left < right {
			list[left] = list[right]
			left++
		}

		for list[left] < mid && left < right {
			left++
		}

		if left < right {
			list[right] = list[left]
			right--
		}
	}

	list[left] = mid

	QuickSort(list[:left])
	QuickSort(list[left+1:])
}

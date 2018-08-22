package sort

func SelectSort(list []int) {
	for i := 0; i < len(list); i++ {
		var exchangeIndex = i
		for j := i + 1; j < len(list); j++ {
			if list[j] < exchangeIndex {
				exchangeIndex = j
			}
		}
		if exchangeIndex != i {
			list[i], list[exchangeIndex] = list[exchangeIndex], list[i]
		}
	}
}

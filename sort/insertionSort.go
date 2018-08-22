package sort

func InsertionSort(list []int) {
	for i := 1; i < len(list); i++ {
		temp := list[i]
		var j int
		for j = i; j > 0 && list[j-1] > temp; j-- {
			list[j] = list[j-1]
		}
		list[j] = temp
	}
}

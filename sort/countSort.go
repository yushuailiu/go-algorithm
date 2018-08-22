package sort

func CountSort(list []int) {
	countList := make(map[int]int)

	for i := 0; i < len(list); i++ {
		countList[list[i]]++
	}
	num := 0
	for key, count := range countList {
		for j := 0; j < count; j++ {
			list[num] = key
			num++
		}
	}
}

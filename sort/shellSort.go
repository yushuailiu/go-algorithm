package sort

func ShellSort(list []int) {
	if len(list) < 2 {
		return
	}
	for step := len(list)/2; step >=1; step /=2 {
		for i := step; i < len(list); i++ {
			temp := list[i]
			var j int
			for j = i; j >= step && list[j - step] >= temp; j -= step {
				list[j] = list[j - step]
			}
			list[j] = temp
		}
	}
}

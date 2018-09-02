// 合并排序 详细解析可以看：https://www.cnblogs.com/chengxiao/p/6194356.html
package sort

func MergeSort(list []int) {
	temp := make([]int, len(list))
	sort(list, 0, len(list)-1, temp)
}

func sort(list []int, start, end int, temp []int) {
	if start < end {
		mid := (end + start) / 2
		sort(list, start, mid, temp)
		sort(list, mid+1, end, temp)
		merge(list, start, mid, end, temp)
	}
}

func merge(list []int, start, mid, end int, temp []int) {
	j := mid + 1
	i := start
	cur := start

	for i <= mid && j <= end {
		if list[i] <= list[j] {
			temp[cur] = list[i]
			i++
		} else {
			temp[cur] = list[j]
			j++
		}
		cur++
	}

	for i <= mid {
		temp[cur] = list[i]
		i++
		cur++
	}

	for j <= end {
		temp[cur] = list[j]
		j++
		cur++
	}

	copy(list[start:end+1], temp[start:end+1])
}

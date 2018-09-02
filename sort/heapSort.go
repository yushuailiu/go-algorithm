// 堆排序，具体讲解可参考 https://www.cnblogs.com/chengxiao/p/6129630.html
package sort

func HeapSort(list []int) {
	for i := len(list)/2 - 1; i >= 0; i-- {
		adjustHeep(list, i)
	}

	for i := len(list) - 1; i > 0; i-- {
		list[0], list[i] = list[i], list[0]
		adjustHeep(list[:i], 0)
	}
}

func adjustHeep(list []int, cur int) {
	temp := list[cur]
	for k := cur*2 + 1; k < len(list); k = k*2 + 1 {
		if k+1 < len(list) && list[k] < list[k+1] {
			k++
		}
		if list[k] > temp {
			list[cur] = list[k]
			cur = k
		} else {
			break
		}
	}
	list[cur] = temp
}

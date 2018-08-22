package search

func BinarySearch(arr []int, find int) int {
	count := len(arr)
	if count == 0 {
		return -1
	}
	low := 0
	high := count - 1

	for low <= high {
		mid := low + (high - low)/2
		if arr[mid] == find {
			return mid
		} else if arr[mid] < find {
			low = mid + 1
		} else if arr[mid] > find {
			high = mid - 1
		}
	}

	return -1
}
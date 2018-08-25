// 查找数组内重复的数字 —— 剑指offer 面试题三。
// 找出数组中重复的数字
package array

import "fmt"

// 查找数组内一个重复的数字（可能有多个重复数字，但是只找一个）。
// 在一个长度为n的数组里的所有数字都在0到n-1的范围内。数组中某些数字是重复的，但是不知道几个数字是重复的，
// 也不知道每个数字重复几次。请找出任意一个重复的数字。
func FindRepeatedNumber(list []int) (bool, int, error) {

	if 0 == len(list) {
		return false, 0, fmt.Errorf("array length can't be 0")
	}

	for _, item := range list {
		if item < 0 || item > len(list)-1 {
			return false, 0, fmt.Errorf("array contain error number")
		}
	}

	for index := range list {
		for list[index] != index {
			if list[index] == list[list[index]] {
				return true, list[index], nil
			}
			list[index], list[list[index]] = list[list[index]], list[index]
		}
	}
	return false, 0, nil
}

// 在一个长度为n+1的数组里的所有的数字都在1-n的范围内，所以数组中至少有一个数字是重复的。请找出数组中任意一个重复数字，但是不能修改数组。
func FindRepeatedNumberWithoutModifyArray(list []int) (int, error) {

	if len(list) <= 1 {
		return -1, fmt.Errorf("array length must bigger then 1")
	}

	for _, item := range list {
		if item <= 0 || item >= len(list) {
			return -1, fmt.Errorf("array contain error number")
		}
	}

	start := 1
	end := len(list) - 1

	for end >= start {
		middle := ((end - start) >> 1) + start

		count := rangeCount(list, start, middle)

		if end == start {
			if count > 1 {
				return start, nil
			} else {
				break
			}
		}

		if count > (middle - start + 1) {
			end = middle
		} else {
			start = middle + 1
		}

	}
	return -1, nil
}

func rangeCount(list []int, start, end int) int {
	var count int

	for _, i := range list {
		if i >= start && i <= end {
			count++
		}
	}
	return count
}

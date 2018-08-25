package divideAndConquer

import (
	"fmt"
	"math"
)

// 该函数用于求数组的其中一个最大子数组。
// 算法是基于分治策略实现的，具体思路可参考书籍 算法导论 第四章的讲解.
func FindMaximumSubArray(list []int) (int, int, int, error) {

	if len(list) == 0 {
		return 0, 0, 0, fmt.Errorf("array length can't be 0")
	}
	if len(list) == 1 {
		return 0, 0, list[0], nil
	}
	left, right, sum := FindMaximumSubArrayByRange(list, 0, len(list)-1)
	return left, right, sum, nil
}

func FindMaximumSubArrayByRange(list []int, low, high int) (int, int, int) {

	if low == high {
		return low, low, list[low]
	}

	mid := (high-low)/2 + low

	leftLow, leftHigh, leftSum := FindMaximumSubArrayByRange(list, low, mid)
	crossLow, crossHigh, crossSum := findMaxCrossSubArray(list, low, mid, high)
	rightLow, rightHigh, rightSum := FindMaximumSubArrayByRange(list, mid+1, high)

	if leftSum >= rightSum && leftSum >= crossSum {
		return leftLow, leftHigh, crossSum
	} else if rightSum >= leftSum && rightSum >= crossSum {
		return rightLow, rightHigh, rightSum
	} else {
		return crossLow, crossHigh, crossSum
	}
}

func findMaxCrossSubArray(list []int, low, mid, high int) (int, int, int) {
	leftMaxSum := math.MinInt32
	leftMaxIndex := mid

	sum := 0
	i := mid
	for i >= low {
		sum += list[i]
		if sum >= leftMaxSum {
			leftMaxIndex = i
			leftMaxSum = sum
		}
		i--
	}

	rightMaxSum := math.MinInt32
	rightMaxIndex := mid
	sum = 0
	i = mid + 1

	for i <= high {
		sum += list[i]
		if sum >= rightMaxSum {
			rightMaxSum = sum
			rightMaxIndex = i
		}
		i++
	}

	return leftMaxIndex, rightMaxIndex, leftMaxSum + rightMaxSum
}

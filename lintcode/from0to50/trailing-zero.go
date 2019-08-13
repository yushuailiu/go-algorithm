package from0to50

// trailingZeros 计算n的阶乘的值的末尾包含0的个数
func trailingZeros (n int64) int64 {
	sum := int64(0)
	for n != 0 {
		n = n/5
		sum += n
	}
	return sum
}

package from0to50

import "time"

// aplusb 不使用+等算数运算符实现两个数相加
func aplusb(a, b int) int {
	if b == 0 {
		return a
	}
	time.Now()
	return aplusb(a ^ b, a & b << 1)
}

package divideAndConquer

import "testing"

var list = []map[string]interface{}{
	{
		"array":  []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7},
		"result": []int{7, 10, 42},
	},
	{
		"array":  []int{},
		"result": []int{0, 0, 0},
	},
}

func TestNormal(t *testing.T) {
	list := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	left, right, sum, err := FindMaximumSubArray(list)

	if err != nil {
		t.Error("test get error:", err)
	}

	if left != 7 || right != 10 || sum != 43 {
		t.Error("test ", list, "should get 7 10 43 but get", left, right, sum)
	}
}

func TestEmptyArray(t *testing.T) {
	left, right, sum, err := FindMaximumSubArray([]int{})

	if left != 0 || right != 0 || sum != 0 || err == nil {
		t.Error("test empty array should get '0 0 0 err' but get'", left, right, sum, err, "'")
	}
}

func TestAllNumberIsPositive(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6}

	left, right, sum, err := FindMaximumSubArray(list)

	if err != nil {
		t.Error("test all number is positive should not get", err)
	}

	if left != 0 || right != 5 || sum != 21 {
		t.Error("test all number is positive should get '0 5 21' but get", left, right, sum)
	}
}

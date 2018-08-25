package array

import (
	"testing"
)

func TestNormalHasRepeatedNumber(t *testing.T) {
	list := []int{2, 3, 1, 0, 2, 5, 3}

	find, repeatedNumber, err := FindRepeatedNumber(list)

	if find == false || repeatedNumber != 2 || err != nil {
		t.Error("test ", list, "should get repeated number 2 but get", repeatedNumber)
	}
}

func TestNormalHasNotRepeatedNumber(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 0}

	find, repeatedNumber, err := FindRepeatedNumber(list)
	if err != nil {
		t.Error("test", list, "not right array")
	}

	if find != false {
		t.Error("test ", list, "should not get repeated number but get repeated number", repeatedNumber)
	}
}

func TestArrayContainWrongNumber(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}

	find, repeatedNumber, err := FindRepeatedNumber(list)

	if find != false || err == nil {
		t.Error("test", list, "should not get repeated number but get", repeatedNumber)
	}
}

func TestWithOutModifyArrayNormal(t *testing.T) {
	list := []int{2, 4, 1, 6, 2, 5, 3}

	repeatedNumber, err := FindRepeatedNumberWithoutModifyArray(list)

	if repeatedNumber != 2 {
		t.Errorf("test %v should get 2 but get %d and err:%v", list, repeatedNumber, err)
	}
}

func TestWithOutModifyArrayNormalWithWrongLength(t *testing.T) {
	list := []int{1}

	repeatedNumber, err := FindRepeatedNumberWithoutModifyArray(list)

	if repeatedNumber != -1 || err == nil {
		t.Errorf("test %v should get -1 but get %d and err:%v", list, repeatedNumber, err)
	}
}
func TestWithOutModifyArrayNormalWithWrongNumber(t *testing.T) {
	list := []int{2, 4, 1, 6, 7, 5, 3}

	repeatedNumber, err := FindRepeatedNumberWithoutModifyArray(list)

	if repeatedNumber != -1 || err == nil {
		t.Errorf("test %v should get -1 but get %d and err:%v", list, repeatedNumber, err)
	}
}

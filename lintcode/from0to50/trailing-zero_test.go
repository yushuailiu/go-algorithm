package from0to50

import (
	"testing"
)

var trailingList = []map[string]int64{
	{
		"n": 5,
	},{
		"n": 2,
	},{
		"n": 20,
	},{
		"n": 11,
	},{
		"n": 11,
	},{
		"n": 9,
	},
}

func TestTrailingZeros(t *testing.T) {
	for _, item := range trailingList {
		num := trailingZeros(item["n"])
		n := item["n"]
		total := int64(1)
		for n > 0  {
			total *= n
			n--
		}

		for num > 0{
			remainder := total%10
			if remainder != 0 {
				t.Errorf("calculate %d trailing zeros error should not be %d", item["n"], remainder)
			}
			num = num/10
		}
	}
}
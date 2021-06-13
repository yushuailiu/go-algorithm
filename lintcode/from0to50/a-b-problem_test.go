package from0to50

import (
	"testing"
)

var list = []map[string]int{
 	{
 		"a": 1,
 		"b": 2,
 		"sum": 3,
	},{
 		"a": -1,
 		"b": 2,
 		"sum": 1,
	},{
 		"a": 22,
 		"b": -10,
 		"sum": 12,
	},{
 		"a": 0,
 		"b": 0,
 		"sum": 0,
	},{
 		"a": 1,
 		"b": 1,
 		"sum": 2,
	},{
 		"a": 2,
 		"b": 0,
 		"sum": 2,
	},
 }

func TestAPlusB(t *testing.T)  {
	for _, item := range list {
		sum := aplusb(item["a"], item["b"])
		if item["sum"] != sum {
			t.Errorf("%d plus %d should get %d but get %d", item["a"], item["b"], item["sum"], sum)
		}
	}
}
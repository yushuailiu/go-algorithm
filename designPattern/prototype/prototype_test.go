package prototype

import (
	"github.com/yushuailiu/assert"
	"testing"
)

func TestPropotye(t *testing.T) {
	manager := NewPrototypeManage()

	circle1 := &Circle{
		Id:   1,
		Type: "circle",
	}
	manager.Set("circle1", circle1)

	circle2 := manager.Get("circle1")
	if circle1 == circle2 {
		t.Error("circle1 should not be circle2")
	}
	assert.Equal(t, circle1, circle2)
}

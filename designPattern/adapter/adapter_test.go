package adapter

import (
	"github.com/yushuailiu/assert"
	"testing"
)

func TestAdapter(t *testing.T) {
	adaptee := NewAdaptee()
	target := NewAdapter(adaptee)
	res := target.Request()

	assert.Equal(t, res, "adaptee specific method")
}

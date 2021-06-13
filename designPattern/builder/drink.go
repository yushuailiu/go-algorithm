package builder

type drink interface {
	Item
}

type Coke struct {
}

func (c *Coke) Name() string {
	return "Coke"
}

func (c *Coke) Price() float64 {
	return 3.4
}

type Coffee struct {
}

func (c *Coffee) Name() string {
	return "Coffee"
}
func (c *Coffee) Price() float64 {
	return 2.5
}

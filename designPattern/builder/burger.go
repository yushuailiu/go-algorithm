package builder

type burger interface {
	Item
}

type VegBurger struct {
}

func (v *VegBurger) Name() string {
	return "VegBurger"
}

func (v *VegBurger) Price() float64 {
	return 4.5
}

type ChickenBurger struct {
}

func (c *ChickenBurger) Name() string {
	return "ChickenBurger"
}

func (c *ChickenBurger) Price() float64 {
	return 5.6
}

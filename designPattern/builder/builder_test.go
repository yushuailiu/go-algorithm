package builder

import "testing"

func TestBuilder(t *testing.T) {
	builder := new(MealBuilder)

	vegBurger := builder.VegMeal()
	chickenBurger := builder.ChickenMeal()

	if vegBurger.Price() != 7.0 {
		t.Error("VegMeal price error")
	}

	if chickenBurger.Price() != 14.6 {
		t.Error("ChickenMeal price error")
	}
}

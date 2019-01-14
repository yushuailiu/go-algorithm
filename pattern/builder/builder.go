package builder

type MealBuilder struct {
}

func (m *MealBuilder) VegMeal() *Meal {
	meal := new(Meal)
	meal.AddItem(new(VegBurger))
	meal.AddItem(new(Coffee))
	return meal
}

func (m *MealBuilder) ChickenMeal() *Meal {
	meal := new(Meal)
	meal.AddItem(new(ChickenBurger))
	meal.AddItem(new(ChickenBurger))
	meal.AddItem(new(Coke))
	return meal
}

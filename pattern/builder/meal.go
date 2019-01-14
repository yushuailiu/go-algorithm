package builder

type Meal struct {
	items []Item
}

func (m *Meal) AddItem(item Item) {
	m.items = append(m.items, item)
}

func (m *Meal) Price() float64 {
	price := 0.0
	for _, i := range m.items {
		price += i.Price()
	}
	return price
}

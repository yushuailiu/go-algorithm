package builder

type Item interface {
	Name() string
	Price() float64
}

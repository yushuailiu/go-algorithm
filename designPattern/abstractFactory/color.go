package abstractFactory

// 颜色接口
type Color interface {
	Color() string
}

// 粉红色主题颜色
type PinkColor struct {
}

func (p *PinkColor) Color() string {
	return "pink"
}

// 蓝色主题颜色
type BlueColor struct {
}

func (b *BlueColor) Color() string {
	return "blue"
}

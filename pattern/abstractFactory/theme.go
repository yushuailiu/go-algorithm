package abstractFactory

// 主题接口
type ThemeFactory interface {
	GetVoice() Voice
	GetColor() Color
}

// 女生主题
type GirlTheme struct {
}

func (g *GirlTheme) GetVoice() Voice {
	return new(GirlVoice)
}

func (g *GirlTheme) GetColor() Color {
	return new(PinkColor)
}

// 男生主题
type BoyTheme struct {
}

func (b *BoyTheme) GetVoice() Voice {
	return new(BoyVoice)
}

func (b *BoyTheme) GetColor() Color {
	return new(BlueColor)
}

package abstractFactory

// 主题接口
type ThemeFactory interface {
	GetVoice() Voice
	GetColor() Color
	GetTheme() *Theme
}

type Theme struct {
	IVoice Voice
	IColor Color
}

func CreateTheme(factory ThemeFactory) *Theme{
	theme := factory.GetTheme()
	theme.IColor = factory.GetColor()
	theme.IVoice = factory.GetVoice()
	return theme
}

// 女生主题
type GirlThemeFactory struct {
}

func (g *GirlThemeFactory) GetVoice() Voice {
	return new(GirlVoice)
}

func (g *GirlThemeFactory) GetColor() Color {
	return new(PinkColor)
}

func (g *GirlThemeFactory) GetTheme() *Theme {
	return new(Theme)
}

// 男生主题
type BoyThemeFactory struct {
}

func (b *BoyThemeFactory) GetVoice() Voice {
	return new(BoyVoice)
}

func (b *BoyThemeFactory) GetColor() Color {
	return new(BlueColor)
}

func (g *BoyThemeFactory) GetTheme() *Theme {
	return new(Theme)
}

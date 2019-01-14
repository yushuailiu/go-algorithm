package abstractFactory

// 声音接口
type Voice interface {
	Voice() string
}

// 女孩声音
type GirlVoice struct {
}

func (g *GirlVoice) Voice() string {
	return "girl"
}

// 男孩声音
type BoyVoice struct {
}

func (b *BoyVoice) Voice() string {
	return "boy"
}

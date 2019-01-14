package abstractFactory

import "testing"

func TestAbstractFactory(t *testing.T) {
	// 实例化女生主题
	girlTheme := new(GirlTheme)
	if girlTheme.GetColor().Color() != "pink" || girlTheme.GetVoice().Voice() != "girl" {
		t.Error("GirlTheme error")
	}
	// 实例化男生主题
	boyTheme := new(BoyTheme)
	if boyTheme.GetColor().Color() != "blue" || boyTheme.GetVoice().Voice() != "boy" {
		t.Error("BoyTheme error")
	}
}

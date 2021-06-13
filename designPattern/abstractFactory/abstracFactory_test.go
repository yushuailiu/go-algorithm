package abstractFactory

import "testing"

func TestAbstractFactory(t *testing.T) {
	// 实例化女生主题
	girlThemeFactory := new(GirlThemeFactory)

	girlTheme := CreateTheme(girlThemeFactory)
	if girlTheme.IColor.Color() != "pink" || girlTheme.IVoice.Voice() != "girl" {
		t.Error("GirlTheme error")
	}
	// 实例化男生主题
	boyThemeFactory := new(BoyThemeFactory)
	boyTheme := CreateTheme(boyThemeFactory)
	if boyTheme.IColor.Color() != "blue" || boyTheme.IVoice.Voice() != "boy" {
		t.Error("BoyTheme error")
	}

}

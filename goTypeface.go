package goTypeface

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
	"runtime"
)

// 声明静态资源,在打包的时候导入到包里面
var (
	//go:embed fonts/simhei.ttf
	NotoSansSC []byte
)

type GoTypeface struct{}

var _ fyne.Theme = (*GoTypeface)(nil)

// Font StaticName 为 fonts 目录下的 ttf 类型的字体文件名
func (m GoTypeface) Font(fyne.TextStyle) fyne.Resource {
	var typeface string = "simhei.ttf"
	sysType := runtime.GOOS
	if sysType == "windows" {
		typeface = "C:/Windows/fonts/msyhbd.ttc"
	}
	//if sysType == "linux" {
	//	//typeface = typeface
	//}

	return &fyne.StaticResource{
		StaticName:    typeface,
		StaticContent: NotoSansSC,
	}
}

func (*GoTypeface) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(n, v)
}

func (*GoTypeface) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (*GoTypeface) Size(n fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(n)
}

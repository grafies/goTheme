

```go

package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"fyneGUI/theme"
)

func main() {
	a := app.New()                          //创建一个fyne应用
	a.Settings().SetTheme(&theme.MyTheme{}) //引入字体,让程序可以显示中文
	w := a.NewWindow("\\(@^0^@)/")          //创建一个窗口的名称
	w.SetContent(widget.NewLabel("hello"))  //窗口内容
	w.ShowAndRun()                          //显示窗口
}

```
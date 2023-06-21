# 使用方式

```sh
go get gitee.com/grafies/goTypeface
```


# 代码演示
```go
package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"gitee.com/grafies/goTypeface"
)

func main() {
	a := app.New()                          //创建一个fyne应用
	a.Settings().SetTheme(&goTypeface.GoTypeface{}) //引入字体,让程序可以显示中文,使用的是黑体
	w := a.NewWindow("\\(@^0^@)/")          //创建一个窗口的名称
	w.SetContent(widget.NewLabel("hello"))  //窗口内容
	w.ShowAndRun()                          //显示窗口
}
```
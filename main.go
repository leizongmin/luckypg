package main

import (
	"strings"

	"github.com/leizongmin/luckypg/internal/console"
)

var header = `
                   _oo8oo_
                  o8888888o
                  88" . "88
                  (| -_- |)
                  0\  =  /0
                ___/'==='\___
              .' \\|     |// '.
             / \\|||  :  |||// \
            / _||||| -:- |||||_ \
           |   | \\\  -  /// |   |
           | \_|  ''\---/''  |_/ |
           \  .-\__  '-'  __/-.  /
         ___'. .'  /--.--\  '. .'___
      ."" '<  '.___\_<|>_/___.'  >' "".
     | | :  ^- \^.:^\ _ /^:.^/ -^  : | |
     \  \ ^-.   \_ __\ /__ _/   .-^ /  /
 =====^-.____^.___ \_____/ ___.^____.-^=====
                   ^=---=^
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
`

var screenWidth = console.TerminalWidth
var boxWidth = 80

func init() {
	header = strings.Trim(header, "\n")
}

func main() {
	L := func(s string) {
		println(s)
	}
	S := func(s string) string {
		return console.CenteredText(s, screenWidth, console.NoColorRenderFunc, console.NoColorRenderFunc)
	}
	L(console.Yellow(S(header)))

	{
		B1 := func(s string) string {
			return console.CenteredText(s, boxWidth, console.NoColorRenderFunc, console.NoColorRenderFunc)
		}
		B2 := func(s string) string {
			return console.FixedWidthText(s, boxWidth-boxWidth/4, console.NoColorRenderFunc, console.XTermColorFunc(0, 250))
		}
		c1 := func(s string) string {
			return console.XTermColor(0, 250, s)
		}
		c2 := func(s string) string {
			return console.XTermColor(28, 250, s)
		}
		c3 := func(s string) string {
			return console.XTermColor(196, 250, s)
		}
		direction := "西南方"
		stars := "★★★★☆"

		lines := []string{
			B2(""),
			B2("今天是2015年12月15日 星期二"),
			B2(""),
			B2("座位朝向：面向" + c2(direction) + c1("写程序，BUG 最少")),
			B2("今日宜饮：绿茶、鲜奶"),
			B2("女神亲近指数：" + c3(stars)),
			B2(""),
		}
		for _, s := range lines {
			L(S(c1(B1(s))))
		}
	}
}

package main

import (
	"strings"

	. "github.com/leizongmin/luckypg/internal/console"
)

var header string = `
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

var screenWidth = TerminalWidth
var boxWidth = 80

func init() {
	header = strings.Trim(header, "\n")
}

func main() {
	println(Yellow(CenteredText(header, screenWidth)))

	{
		c1 := func (s string) string{
			return XTermColor(0, 250, s)
		}
		c2 := func (s string) string{
			return XTermColor(28, 250, s)
		}
		c3 := func (s string) string{
			return XTermColor(196, 250, s)
		}
		println(CenteredText(CenteredText("", boxWidth), screenWidth))
		println(CenteredText(FixedWidthText("今天是2015年12月15日 星期二", boxWidth), screenWidth))
		println(CenteredText(FixedWidthText("", boxWidth), screenWidth))
		println(CenteredText(FixedWidthText("座位朝向：面向"+c2( "西南方")+"写程序，BUG 最少", boxWidth), screenWidth))
		println(CenteredText(FixedWidthText("今日宜饮：绿茶、鲜奶", boxWidth), screenWidth))
		println(CenteredText(FixedWidthText("女神亲近指数："+c3("★★★★☆"), boxWidth), screenWidth))
		println(CenteredText(CenteredText("", boxWidth), screenWidth))
	}
}

func color0(s string) string {
	return XTermColor(227, 0, s)
}

func color1(s string) string {
	return XTermColor(0, 250, s)
}

func color2(s string) string {
	return XTermColor(0, 229, s)
}

func color3(s string) string {
	return XTermColor(0, 220, s)
}

func color4(s string) string {
	return XTermColor(0, 210, s)
}

func color5(s string) string {
	return XTermColor(0, 196, s)
}

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
	L("")

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
	{
		headerBoxWidth := 4
		bodyBoxWidth := boxWidth/2 - headerBoxWidth
		B := func(s string) string {
			return console.CenteredText(s, boxWidth, console.NoColorRenderFunc, console.NoColorRenderFunc)
		}
		B1 := func(s string) string {
			c := console.XTermColorFunc(0, 220)
			return console.CenteredText(s, headerBoxWidth, c, c)
		}
		B2 := func(s string) string {
			c := console.XTermColorFunc(0, 229)
			return console.FixedWidthText(s, bodyBoxWidth, c, c)
		}
		B3 := func(s string) string {
			c := console.XTermColorFunc(0, 196)
			return console.CenteredText(s, headerBoxWidth, c, c)
		}
		B4 := func(s string) string {
			c := console.XTermColorFunc(0, 210)
			return console.FixedWidthText(s, bodyBoxWidth, c, c)
		}
		goodLeftLines := []string{
			B1(""),
			B1(""),
			B1(""),
			B1(""),
			B1(""),
			B1(" 宜 "),
			B1(""),
			B1(""),
			B1(""),
			B1(""),
		}
		goodRightLines := []string{
			B2(""),
			B2("  白天上线"),
			B2("    今天白天上线是安全的"),
			B2("  招人"),
			B2("    你面前这位有成为牛人的潜质"),
			B2("  提交代码"),
			B2("    遇到冲突的几率是最低的"),
			B2("  代码审核"),
			B2("    发现重要问题的几率大大增加"),
			B2(""),
		}
		badLeftLines := []string{
			B3(""),
			B3(""),
			B3(""),
			B3(""),
			B3(""),
			B3(" 忌 "),
			B3(""),
			B3(""),
			B3(""),
			B3(""),
		}
		badRightLines := []string{
			B4(""),
			B4("  开会"),
			B4("    小心被扣屎盆子背黑锅"),
			B4("  晚上上线"),
			B4("    你白天已经精疲力尽了"),
			B4("  修复BUG"),
			B4("    新产生的BUG将比修复的更多"),
			B4("  需求审核"),
			B4("    未知"),
			B4(""),
		}
		for i := range goodLeftLines {
			L(S(B(goodLeftLines[i] + goodRightLines[i] + badLeftLines[i] + badRightLines[i])))
		}
	}
	L("")
}

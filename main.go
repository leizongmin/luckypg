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
	L("")
	{
		B := func(s string) string {
			c := console.XTermColorFunc(245, 250)
			return console.CenteredText(s, boxWidth, c, c)
		}
		L(S(B(header + "\n")))
	}
	{
		B1 := func(s string) string {
			c := console.XTermColorFunc(0, 250)
			return console.CenteredText(s, boxWidth, c, c)
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

		date := "2015年12月15日"
		week := "星期二"
		direction := "西南方"
		stars := "★★★★☆"
		drink := "绿茶、鲜奶"

		lines := []string{
			B2(""),
			B2("今天是" + date + " " + week),
			B2(""),
			B2("座位朝向：面向" + c2(direction) + c1("写程序，BUG 最少")),
			B2("今日宜饮：" + drink),
			B2("女神亲近指数：" + c3(stars)),
			B2(""),
		}
		for _, s := range lines {
			L(S(B1(s)))
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
		B21 := func(s string) string {
			c := console.XTermColorFunc(0, 229)
			return console.FixedWidthText(s, bodyBoxWidth, c, c)
		}
		B22 := func(s string) string {
			c := console.XTermColorFunc(240, 229)
			return console.FixedWidthText(s, bodyBoxWidth, c, c)
		}
		B3 := func(s string) string {
			c := console.XTermColorFunc(0, 196)
			return console.CenteredText(s, headerBoxWidth, c, c)
		}
		B41 := func(s string) string {
			c := console.XTermColorFunc(0, 210)
			return console.FixedWidthText(s, bodyBoxWidth, c, c)
		}
		B42 := func(s string) string {
			c := console.XTermColorFunc(240, 210)
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
			B21(""),
			B21("  白天上线"),
			B22("    今天白天上线是安全的"),
			B21("  招人"),
			B22("    你面前这位有成为牛人的潜质"),
			B21("  提交代码"),
			B22("    遇到冲突的几率是最低的"),
			B21("  代码审核"),
			B22("    发现重要问题的几率大大增加"),
			B21(""),
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
			B41(""),
			B41("  开会"),
			B42("    小心被扣屎盆子背黑锅"),
			B41("  晚上上线"),
			B42("    你白天已经精疲力尽了"),
			B41("  修复BUG"),
			B42("    新产生的BUG将比修复的更多"),
			B41("  需求审核"),
			B42("    未知"),
			B41(""),
		}
		for i := range goodLeftLines {
			L(S(B(goodLeftLines[i] + goodRightLines[i] + badLeftLines[i] + badRightLines[i])))
		}
	}
	L("")
}

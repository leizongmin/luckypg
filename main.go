package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/leizongmin/luckypg/internal/console"
	"github.com/leizongmin/luckypg/internal/data"
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
	today := time.Now()
	L := func(s string) {
		fmt.Println(s)
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

		date := data.TodayDate(today)
		direction := data.TodayDirection(today)
		stars := data.TodayStars(today)
		drink := data.TodayDrink(today)

		lines := []string{
			B2(""),
			B2("今天是" + date),
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

		luck := data.TodayLuck(today, 4)

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
			B21("  " + luck.Good[0].Name),
			B22("    " + luck.Good[0].Good),
			B21("  " + luck.Good[1].Name),
			B22("    " + luck.Good[1].Good),
			B21("  " + luck.Good[2].Name),
			B22("    " + luck.Good[2].Good),
			B21("  " + luck.Good[3].Name),
			B22("    " + luck.Good[3].Good),
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
			B41("  " + luck.Bad[0].Name),
			B42("    " + luck.Bad[0].Bad),
			B41("  " + luck.Bad[1].Name),
			B42("    " + luck.Bad[1].Bad),
			B41("  " + luck.Bad[2].Name),
			B42("    " + luck.Bad[2].Bad),
			B41("  " + luck.Bad[3].Name),
			B42("    " + luck.Bad[3].Bad),
			B41(""),
		}
		for i := range goodLeftLines {
			L(S(B(goodLeftLines[i] + goodRightLines[i] + badLeftLines[i] + badRightLines[i])))
		}
	}
	L("")
}

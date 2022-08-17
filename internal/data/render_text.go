package data

import (
	"github.com/leizongmin/luckypg/internal/nostd"
)

// RenderText 解析占位符并替换成随机内容
// %v = 随机变量名称
// %t = 随机工具名称
// %l = 随机整数
func RenderText(today nostd.Time, s string) string {
	dayNum := TimeToDateNumber(today)

	if nostd.Contains(s, "%v") {
		s = nostd.ReplaceAll(s, "%v", DefineVariableNames[RandomNumber(dayNum, 12)%len(DefineVariableNames)])
	}

	if nostd.Contains(s, "%t") {
		s = nostd.ReplaceAll(s, "%t", DefineTools[RandomNumber(dayNum, 11)%len(DefineTools)])
	}

	if nostd.Contains(s, "%l") {
		s = nostd.ReplaceAll(s, "%l", nostd.Itoa(RandomNumber(dayNum, 12)%247+30))
	}

	return s
}

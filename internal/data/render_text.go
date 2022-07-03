package data

import (
	"github.com/leizongmin/luckypg/internal/nostd"
)

// RenderText 解析占位符并替换成随机内容
// %v = 随机变量名称
// %t = 随机工具名称
// %l = 随机整数
func RenderText(today nostd.Time, event ActivityItem) ActivityItem {
	result := ActivityItem{
		Name: event.Name,
		Good: event.Good,
		Bad:  event.Bad,
	}
	dayNum := TimeToDateNumber(today)
	if nostd.Contains(result.Name, "%v") {
		s := DefineVariableNames[RandomNumber(dayNum, 12)%len(DefineVariableNames)]
		result.Name = nostd.ReplaceAll(result.Name, "%v", s)
	}

	if nostd.Contains(result.Name, "%t") {
		s := DefineTools[RandomNumber(dayNum, 11)%len(DefineTools)]
		result.Name = nostd.ReplaceAll(result.Name, "%t", s)
	}

	if nostd.Contains(result.Name, "%l") {
		s := nostd.Itoa(RandomNumber(dayNum, 12)%247 + 30)
		result.Name = nostd.ReplaceAll(result.Name, "%l", s)
	}

	return result
}

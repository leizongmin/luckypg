package data

import (
	"strconv"
	"strings"
	"time"
)

// RenderText 解析占位符并替换成随机内容
// %v = 随机变量名称
// %t = 随机工具名称
// %l = 随机整数
func RenderText(today time.Time, event ActivityItem) ActivityItem {
	result := ActivityItem{
		Name: event.Name,
		Good: event.Good,
		Bad:  event.Bad,
	}
	dayNum := TimeToDateNumber(today)
	if strings.Contains(result.Name, "%v") {
		s := DefineVariableNames[RandomNumber(dayNum, 12)%len(DefineVariableNames)]
		result.Name = strings.ReplaceAll(result.Name, "%v", s)
	}

	if strings.Contains(result.Name, "%t") {
		s := DefineTools[RandomNumber(dayNum, 11)%len(DefineTools)]
		result.Name = strings.ReplaceAll(result.Name, "%t", s)
	}

	if strings.Contains(result.Name, "%l") {
		s := strconv.Itoa(RandomNumber(dayNum, 12)%247 + 30)
		result.Name = strings.ReplaceAll(result.Name, "%l", s)
	}

	return result
}

package data

import (
	"strings"

	"github.com/leizongmin/luckypg/internal/nostd"
)

// RandomNumber 注意：本程序中的“随机”都是伪随机概念，以当前的天为种子。
func RandomNumber(daySeed int, indexSeed int) int {
	n := daySeed % 11117
	for i := 0; i < 100+indexSeed; i++ {
		n = n * n
		n = n % 11117 // 11117 是个质数
	}
	return n
}

// PickRandomActivityItems 从数组中随机挑选 size 个
func PickRandomActivityItems(today nostd.Time, array []ActivityItem, size int) []ActivityItem {
	result := array[:]
	daySeed := TimeToDateNumber(today)
	for j := 0; j < len(array)-size; j++ {
		i := RandomNumber(daySeed, j) % len(result)
		result = append(result[0:i], result[i+1:]...)
	}
	return result
}

// PickRandomStringItems 从数组中随机挑选 size 个
func PickRandomStringItems(today nostd.Time, array []string, size int) []string {
	result := array[:]
	daySeed := TimeToDateNumber(today)
	for j := 0; j < len(array)-size; j++ {
		i := RandomNumber(daySeed, j) % len(result)
		result = append(result[0:i], result[i+1:]...)
	}
	return result
}

// PickRandomActivity 从 activities 中随机挑选 size 个
func PickRandomActivity(today nostd.Time, list []ActivityItem, size int) []ActivityItem {
	pickedEvents := PickRandomActivityItems(today, list, size)
	for i, item := range pickedEvents {
		pickedEvents[i] = RenderText(today, item)
	}
	return pickedEvents
}

// WeekdayString 返回星期几的文本描述
func WeekdayString(weekday nostd.Weekday) string {
	n := int(weekday)
	if n < 0 || n > 6 {
		return "*"
	}
	return DefineWeeks[n]
}

// TodayDate 今天的日期
func TodayDate(today nostd.Time) string {
	y := nostd.Itoa(today.Year())
	m := nostd.Itoa(int(today.Month()))
	d := nostd.Itoa(today.Day())
	w := WeekdayString(today.Weekday())
	return "今天是" + y + "年" + m + "月" + d + "日 星期" + w
}

// GenerateStars 生成星星
func GenerateStars(num int, max int) string {
	s := ""
	for i := 0; i < num; i++ {
		s += "★"
	}
	for i := num; i < max; i++ {
		s += "☆"
	}
	return s
}

// TimeToDateNumber 将日期转换为 int 描述，如 20120708
func TimeToDateNumber(today nostd.Time) int {
	return today.Year()*10000 + int(today.Month())*100 + today.Day()
}

// IsWeekend 是否为周末
func IsWeekend(today nostd.Time) bool {
	return today.Day() == 0 || today.Day() == 6
}

// GetActivities 去掉一些不合今日的事件
func GetActivities(today nostd.Time) []ActivityItem {
	// 周末的话，只留下 weekend = true 的事件
	if IsWeekend(today) {
		result := make([]ActivityItem, 0)
		for _, item := range DefineActivities {
			if item.Weekend {
				result = append(result, item)
			}
		}
		return result
	}
	return DefineActivities[:]
}

type ActivityGoodBad struct {
	Good []ActivityItem
	Bad  []ActivityItem
}

// GetSpecials 今天的特别事件
func GetSpecials(today nostd.Time) ActivityGoodBad {
	goodList := make([]ActivityItem, 0)
	badList := make([]ActivityItem, 0)
	daySeed := TimeToDateNumber(today)
	for _, special := range DefineSpecials {
		if daySeed == special.Date || daySeed-today.Year()*10000 == special.Date {
			if special.Type == "good" {
				goodList = append(goodList, ActivityItem{
					Name: special.Name,
					Good: special.Description,
				})
			} else {
				badList = append(badList, ActivityItem{
					Name: special.Name,
					Bad:  special.Description,
				})
			}
		}
	}
	return ActivityGoodBad{Good: goodList, Bad: badList}
}

// TodayLuck 生成今日运势
func TodayLuck(today nostd.Time, n int) ActivityGoodBad {
	// daySeed := TimeToDateNumber(today)
	numGood := n // RandomNumber(daySeed, 98)%3 + 2
	numBad := n  // RandomNumber(daySeed, 87)%3 + 2
	activities := PickRandomActivity(today, GetActivities(today), numGood+numBad)

	special := GetSpecials(today)
	goodList := special.Good
	badList := special.Bad

	for i := 0; i < numGood; i++ {
		goodList = append(goodList, activities[i])
	}
	for i := 0; i < numBad; i++ {
		badList = append(badList, activities[numGood+i])
	}

	return ActivityGoodBad{Good: goodList[:n], Bad: badList[:n]}
}

// TodayDirection 获得座位朝向
func TodayDirection(today nostd.Time) string {
	return DefineDirections[RandomNumber(TimeToDateNumber(today), 2)%len(DefineDirections)]
}

// TodayDrink 获得今日宜饮
func TodayDrink(today nostd.Time) string {
	list := PickRandomStringItems(today, DefineDrinks, 2)
	return strings.Join(list, "，")
}

// TodayStars 获得女神亲近指数
func TodayStars(today nostd.Time) string {
	return GenerateStars(RandomNumber(TimeToDateNumber(today), 6)%5+1, 5)
}

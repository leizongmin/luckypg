package internal

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var weeks = []string{"日", "一", "二", "三", "四", "五", "六"}
var directions = []string{"北方", "东北方", "东方", "东南方", "南方", "西南方", "西方", "西北方"}
var tools = []string{"Eclipse写程序", "MSOffice写文档", "记事本写程序", "Windows8", "Linux", "MacOS", "IE", "Android设备", "iOS设备"}
var varNames = []string{"jieguo", "huodong", "pay", "expire", "zhangdan", "every", "free", "i1", "a", "virtual", "ad", "spider", "mima", "pass", "ui"}
var drinks = []string{"水", "茶", "红茶", "绿茶", "咖啡", "奶茶", "可乐", "鲜奶", "豆奶", "果汁", "果味汽水", "苏打水", "运动饮料", "酸奶", "酒"}

type activityItem struct {
	name    string
	good    string
	bad     string
	weekend bool
}

var activities = []activityItem{
	{
		name: "写单元测试",
		good: "写单元测试将减少出错",
		bad:  "写单元测试会降低你的开发效率",
	}, {
		name:    "洗澡",
		good:    "你几天没洗澡了？",
		bad:     "会把设计方面的灵感洗掉",
		weekend: true,
	}, {
		name:    "锻炼一下身体",
		good:    "",
		bad:     "能量没消耗多少，吃得却更多",
		weekend: true,
	}, {
		name:    "抽烟",
		good:    "抽烟有利于提神，增加思维敏捷",
		bad:     "除非你活够了，死得早点没关系",
		weekend: true,
	}, {
		name: "白天上线",
		good: "今天白天上线是安全的",
		bad:  "可能导致灾难性后果",
	}, {
		name: "重构",
		good: "代码质量得到提高",
		bad:  "你很有可能会陷入泥潭",
	}, {
		name: "使用%t",
		good: "你看起来更有品位",
		bad:  "别人会觉得你在装逼",
	}, {
		name: "跳槽",
		good: "该放手时就放手",
		bad:  "鉴于当前的经济形势，你的下一份工作未必比现在强",
	}, {
		name: "招人",
		good: "你面前这位有成为牛人的潜质",
		bad:  "这人会写程序吗？",
	}, {
		name: "面试",
		good: "面试官今天心情很好",
		bad:  "面试官不爽，会拿你出气",
	}, {
		name: "提交辞职申请",
		good: "公司找到了一个比你更能干更便宜的家伙，巴不得你赶快滚蛋",
		bad:  "鉴于当前的经济形势，你的下一份工作未必比现在强",
	}, {
		name: "申请加薪",
		good: "老板今天心情很好",
		bad:  "公司正在考虑裁员",
	}, {
		name:    "晚上加班",
		good:    "晚上是程序员精神最好的时候",
		bad:     "",
		weekend: true,
	}, {
		name:    "在妹子面前吹牛",
		good:    "改善你矮穷挫的形象",
		bad:     "会被识破",
		weekend: true,
	}, {
		name:    "撸管",
		good:    "避免缓冲区溢出",
		bad:     "强撸灰飞烟灭",
		weekend: true,
	}, {
		name:    "浏览成人网站",
		good:    "重拾对生活的信心",
		bad:     "你会心神不宁",
		weekend: true,
	}, {
		name: "命名变量\"%v\"",
		good: "",
		bad:  "",
	}, {
		name: "写超过%l行的方法",
		good: "你的代码组织的很好，长一点没关系",
		bad:  "你的代码将混乱不堪，你自己都看不懂",
	}, {
		name: "提交代码",
		good: "遇到冲突的几率是最低的",
		bad:  "你遇到的一大堆冲突会让你觉得自己是不是时间穿越了",
	}, {
		name: "代码复审",
		good: "发现重要问题的几率大大增加",
		bad:  "你什么问题都发现不了，白白浪费时间",
	}, {
		name: "开会",
		good: "写代码之余放松一下打个盹，有益健康",
		bad:  "小心被扣屎盆子背黑锅",
	}, {
		name:    "打DOTA",
		good:    "你将有如神助",
		bad:     "你会被虐的很惨",
		weekend: true,
	}, {
		name: "晚上上线",
		good: "晚上是程序员精神最好的时候",
		bad:  "你白天已经筋疲力尽了",
	}, {
		name: "修复BUG",
		good: "你今天对BUG的嗅觉大大提高",
		bad:  "新产生的BUG将比修复的更多",
	}, {
		name: "设计评审",
		good: "设计评审会议将变成头脑风暴",
		bad:  "人人筋疲力尽，评审就这么过了",
	}, {
		name: "需求评审",
		good: "",
		bad:  "",
	}, {
		name:    "上微博",
		good:    "今天发生的事不能错过",
		bad:     "今天的微博充满负能量",
		weekend: true,
	}, {
		name:    "上AB站",
		good:    "还需要理由吗？",
		bad:     "满屏兄贵亮瞎你的眼",
		weekend: true,
	}, {
		name:    "玩FlappyBird",
		good:    "今天破纪录的几率很高",
		bad:     "除非你想玩到把手机砸了",
		weekend: true,
	},
}

type specialItem struct {
	date        int
	typ         string
	name        string
	description string
}

var specials = []specialItem{
	{
		date:        20151224,
		typ:         "good",
		name:        "待在男（女）友身边",
		description: "今天晚上不要加班了吧？",
	}, {
		date:        20160214,
		typ:         "bad",
		name:        "待在男（女）友身边",
		description: "脱团火葬场，入团保平安。",
	},
}

/// 注意：本程序中的“随机”都是伪随机概念，以当前的天为种子。
func random(daySeed int, indexSeed int) int {
	n := daySeed % 11117
	for i := 0; i < 100+indexSeed; i++ {
		n = n * n
		n = n % 11117 // 11117 是个质数
	}
	return n
}

/// 从数组中随机挑选 size 个
func pickRandomActivityItems(today time.Time, array []activityItem, size int) []activityItem {
	result := array[:]
	dayNum := getDayNumber(today)
	for j := 0; j < len(array)-size; j++ {
		i := random(dayNum, j) % len(result)
		result = append(result[0:i], result[i+1:]...)
	}
	return result
}

/// 从数组中随机挑选 size 个
func pickRandomStringItems(today time.Time, array []string, size int) []string {
	result := array[:]
	dayNum := getDayNumber(today)
	for j := 0; j < len(array)-size; j++ {
		i := random(dayNum, j) % len(result)
		result = append(result[0:i], result[i+1:]...)
	}
	return result
}

// 从 activities 中随机挑选 size 个
func pickRandomActivity(today time.Time, list []activityItem, size int) []activityItem {
	pickedEvents := pickRandomActivityItems(today, list, size)
	for i, item := range pickedEvents {
		pickedEvents[i] = parse(today, item)
	}
	return pickedEvents
}

func getWeekdayString(weekday time.Weekday) string {
	n := int(weekday)
	if n < 0 || n > 6 {
		return "*"
	}
	return weeks[n]
}

func getTodayString(today time.Time) string {
	return fmt.Sprintf("今天是%d年%d月%d日 星期%s", today.Year(), today.Month(), today.Day(), getWeekdayString(today.Weekday()))
}

func generateStar(num int, max int) string {
	s := ""
	for i := 0; i < num; i++ {
		s += "★"
	}
	for i := num; i < max; i++ {
		s += "☆"
	}
	return s
}

func getDayNumber(today time.Time) int {
	return today.Year()*10000 + int(today.Month())*100 + today.Day()
}

func isWeekend(today time.Time) bool {
	return today.Day() == 0 || today.Day() == 6
}

/// 去掉一些不合今日的事件
func getFilteredActivities(today time.Time) []activityItem {
	// 周末的话，只留下 weekend = true 的事件
	if isWeekend(today) {
		result = make([]activityItem, 0)
		for _, item := range activities {
			if item.weekend {
				result = append(result, item)
			}
		}
		return result
	}

	return activities[:]
}

type specialsGoodBad struct {
	good []specialItem
	bad  []specialItem
}

type activityGoodBad struct {
	good []activityItem
	bad  []activityItem
}

/// 添加预定义事件
func pickSpecials() specialsGoodBad {
	goodList := make([]specialItem, 0)
	badList := make([]specialItem, 0)
	dayNum := getDayNumber(today)
	for _, special := range specials {
		if dayNum == special.date {
			if special.typ == "good" {
				goodList = append(goodList, specialItem{
					name: special.name,
					good: special.description,
				})
			} else {
				badList = append(badList, specialItem{
					name: special.name,
					bad:  special.description,
				})
			}
		}
	}

	return specialsGoodBad{good: goodList, bad: badList}
}

/// 生成今日运势
func pickTodaysLuck(today time.Time) specialsGoodBad {
	dayNum := getDayNumber(today)
	filteredActivities := getFilteredActivities(today)

	numGood := random(dayNum, 98)%3 + 2
	numBad := random(dayNum, 87)%3 + 2
	eventArr := pickRandomActivity(today, filteredActivities, numGood+numBad)

	special := pickSpecials()

	goodList := special.good
	badList := special.bad

	for i := 0; i < numGood; i++ {
		goodList = append(goodList, eventArr[i])
	}

	for i := 0; i < numBad; i++ {
		badList = append(badList, eventArr[numGood+i])
	}

	return specialsGoodBad{good: goodList[:4], bad: badList[:4]}
}

/// 获得座位朝向
func getDirectionString(today time.Time) string {
	return directions[random(getDayNumber(today), 2)%len(directions)]
}

/// 获得今日宜饮
func getDrinkString(today time.Time) string {
	list := pickRandomStringItems(today, drinks, 2)
	return strings.Join(list, "，")
}

/// 获得女神亲近指数
func getStarString(today time.Time) string {
	return generateStar(random(getDayNumber(today), 6)%5+1, 5)
}

// 解析占位符并替换成随机内容
func parse(today time.Time, event activityItem) activityItem {
	result := activityItem{
		name: event.name,
		good: event.good,
		bad:  event.bad,
	}
	dayNum := getDayNumber(today)
	if strings.Contains(result.name, "%v") {
		s := varNames[random(dayNum, 12)%len(varNames)]
		result.name = strings.ReplaceAll(result.name, "%v", s)
	}

	if strings.Contains(result.name, "%t") {
		s := tools[random(dayNum, 11)%len(tools)]
		result.name = strings.ReplaceAll(result.name, "%t", s)
	}

	if strings.Contains(result.name, "%l") {
		s := strconv.Itoa(random(dayNum, 12)%247 + 30)
		result.name = strings.ReplaceAll(result.name, "%l", s)
	}

	return result
}

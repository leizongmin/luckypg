package data

var DefineWeeks = []string{"日", "一", "二", "三", "四", "五", "六"}
var DefineDirections = []string{"北方", "东北方", "东方", "东南方", "南方", "西南方", "西方", "西北方"}
var DefineTools = []string{"Eclipse写程序", "MSOffice写文档", "记事本写程序", "Windows8", "Linux", "MacOS", "IE", "Android设备", "iOS设备"}
var DefineVariableNames = []string{"jieguo", "huodong", "pay", "expire", "zhangdan", "every", "free", "i1", "a", "virtual", "ad", "spider", "mima", "pass", "ui"}
var DefineDrinks = []string{"水", "茶", "红茶", "绿茶", "咖啡", "奶茶", "可乐", "鲜奶", "豆奶", "果汁", "果味汽水", "苏打水", "运动饮料", "酸奶", "酒"}

// ActivityItem 活动
type ActivityItem struct {
	Name    string // 类目
	Good    string // 当为【宜】时的解释
	Bad     string // 当为【忌】时当解释
	Weekend bool   // 是否限定只周末出现
}

var DefineActivities = []ActivityItem{
	{
		Name: "写单元测试",
		Good: "写单元测试将减少出错",
		Bad:  "写单元测试会降低你的开发效率",
	}, {
		Name:    "洗澡",
		Good:    "你几天没洗澡了？",
		Bad:     "会把设计方面的灵感洗掉",
		Weekend: true,
	}, {
		Name:    "锻炼一下身体",
		Good:    "",
		Bad:     "能量没消耗多少，吃得却更多",
		Weekend: true,
	}, {
		Name:    "抽烟",
		Good:    "抽烟有利于提神，增加思维敏捷",
		Bad:     "除非你活够了，死得早点没关系",
		Weekend: true,
	}, {
		Name: "白天上线",
		Good: "今天白天上线是安全的",
		Bad:  "可能导致灾难性后果",
	}, {
		Name: "重构",
		Good: "代码质量得到提高",
		Bad:  "你很有可能会陷入泥潭",
	}, {
		Name: "使用%t",
		Good: "你看起来更有品位",
		Bad:  "别人会觉得你在装逼",
	}, {
		Name: "跳槽",
		Good: "该放手时就放手",
		Bad:  "鉴于当前的经济形势，你的下一份工作未必比现在强",
	}, {
		Name: "招人",
		Good: "你面前这位有成为牛人的潜质",
		Bad:  "这人会写程序吗？",
	}, {
		Name: "面试",
		Good: "面试官今天心情很好",
		Bad:  "面试官不爽，会拿你出气",
	}, {
		Name: "提交辞职申请",
		Good: "公司找到了一个比你更能干更便宜的家伙，巴不得你赶快滚蛋",
		Bad:  "鉴于当前的经济形势，你的下一份工作未必比现在强",
	}, {
		Name: "申请加薪",
		Good: "老板今天心情很好",
		Bad:  "公司正在考虑裁员",
	}, {
		Name:    "晚上加班",
		Good:    "晚上是程序员精神最好的时候",
		Bad:     "",
		Weekend: true,
	}, {
		Name:    "在妹子面前吹牛",
		Good:    "改善你矮穷挫的形象",
		Bad:     "会被识破",
		Weekend: true,
	}, {
		Name:    "撸管",
		Good:    "避免缓冲区溢出",
		Bad:     "强撸灰飞烟灭",
		Weekend: true,
	}, {
		Name:    "浏览成人网站",
		Good:    "重拾对生活的信心",
		Bad:     "你会心神不宁",
		Weekend: true,
	}, {
		Name: "命名变量\"%v\"",
		Good: "",
		Bad:  "",
	}, {
		Name: "写超过%l行的方法",
		Good: "你的代码组织的很好，长一点没关系",
		Bad:  "你的代码将混乱不堪，你自己都看不懂",
	}, {
		Name: "提交代码",
		Good: "遇到冲突的几率是最低的",
		Bad:  "你遇到的一大堆冲突会让你觉得自己是不是时间穿越了",
	}, {
		Name: "代码复审",
		Good: "发现重要问题的几率大大增加",
		Bad:  "你什么问题都发现不了，白白浪费时间",
	}, {
		Name: "开会",
		Good: "写代码之余放松一下打个盹，有益健康",
		Bad:  "小心被扣屎盆子背黑锅",
	}, {
		Name:    "打DOTA",
		Good:    "你将有如神助",
		Bad:     "你会被虐的很惨",
		Weekend: true,
	}, {
		Name: "晚上上线",
		Good: "晚上是程序员精神最好的时候",
		Bad:  "你白天已经筋疲力尽了",
	}, {
		Name: "修复BUG",
		Good: "你今天对BUG的嗅觉大大提高",
		Bad:  "新产生的BUG将比修复的更多",
	}, {
		Name: "设计评审",
		Good: "设计评审会议将变成头脑风暴",
		Bad:  "人人筋疲力尽，评审就这么过了",
	}, {
		Name: "需求评审",
		Good: "",
		Bad:  "",
	}, {
		Name:    "上微博",
		Good:    "今天发生的事不能错过",
		Bad:     "今天的微博充满负能量",
		Weekend: true,
	}, {
		Name:    "上AB站",
		Good:    "还需要理由吗？",
		Bad:     "满屏兄贵亮瞎你的眼",
		Weekend: true,
	}, {
		Name:    "玩FlappyBird",
		Good:    "今天破纪录的几率很高",
		Bad:     "除非你想玩到把手机砸了",
		Weekend: true,
	},
}

// SpecialItem 特别日子的活动
type SpecialItem struct {
	Date        int    // 日期，如 20120708
	Type        string // 类型，good 或 bad
	Name        string // 类目
	Description string // 说明
}

var DefineSpecials = []SpecialItem{
	{
		Date:        20151224,
		Type:        "good",
		Name:        "待在男（女）友身边",
		Description: "今天晚上不要加班了吧？",
	}, {
		Date:        20160214,
		Type:        "bad",
		Name:        "待在男（女）友身边",
		Description: "脱团火葬场，入团保平安。",
	},
}

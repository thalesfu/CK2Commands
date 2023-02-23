package people

import "bufio"

var badTraits = []string{
	"drunkard",        // 酒鬼
	"infirm",          // 体弱多病
	"possessed",       // 恶魔附身
	"is_fat",          // 肥胖
	"is_malnourished", // 营养不良
	"clubfooted",      // 足内翻
	"dwarf",           // 侏儒
	"harelip",         // 兔唇
	"hunchback",       // 驼背
	"imbecile",        // 痴呆
	"inbred",          // 近亲缺陷
	"lisp",            // 口齿不清
	"slow",            // 迟钝
	"stutter",         // 口吃
	"ugly",            // 丑陋
	"weak",            // 虚弱
	"giant",           // 巨人症
	"lefthanded",      // 左撇子
	"feeble",          // 脆弱
	"sturdy",          // 结实
	"dull",            // 迟钝
	"uncouth",         // 邋遢
	"gluttonous",      // 暴食
	"greedy",          // 贪婪
	"slothful",        // 懒惰
	"wroth",           // 易怒
	"envious",         // 嫉妒
	"proud",           // 骄傲
	"chaste",          // 忠贞
	"ambitious",       // 野心勃勃
	"arbitrary",       // 专断
	"craven",          // 怯懦
	"cruel",           // 残暴
	"cynical",         // 愤世嫉俗
	"deceitful",       // 狡诈
	"paranoid",        // 多疑
	"shy",             // 害羞
	"stubborn",        // 固执
	"trusting",        // 轻信他人
	"homosexual",      // 同性恋
	"celibate",        // 独身主义
}

func RemoveBadTraits(writer *bufio.Writer, peopleId int) {
	for _, trait := range badTraits {
		removeTrait(writer, trait, peopleId)
	}

	cureIlls(writer, peopleId)
	cureWound(writer, peopleId)
}

func RemoveBad(peopleId ...int) {
	var functions = []buildFunc{
		cureIlls,
		cureWound,
	}

	buildPeople("removebad", functions, peopleId...)
}

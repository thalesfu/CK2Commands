package people

import "bufio"

var commonGoodTraits = []string{
	"shrewd",     // 精明
	"fair",       // 魅力非凡
	"genius",     // 天才
	"strong",     // 强壮
	"robust",     // 健壮
	"groomed",    // 整洁
	"lustful",    // 色欲
	"temperate",  // 温和
	"charitable", // 慷慨
	"diligent",   // 勤奋
	"patient",    // 耐心
	"kind",       // 仁慈
	"humble",     // 谦卑
}

func addCommonGoodTraits(writer *bufio.Writer, peopleId int) {
	for _, trait := range commonGoodTraits {
		writeAddTrait(writer, trait, peopleId)
	}
}

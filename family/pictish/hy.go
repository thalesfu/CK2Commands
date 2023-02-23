package pictish

import (
	"github.com/thalesfu/CK2Commands/people"
)

func Hy() {
	people.Pollinate(getPeoples(), "pictish")
}

func getPeoples() []people.Couple {
	var couples []people.Couple
	couples = append(couples, people.Couple{Husband: 2794096, Wife: 2747106}) // 乌苏尔特 西诺德.韦尔赫敏礼.乌伊布
	couples = append(couples, people.Couple{Husband: 2783181, Wife: 2795786}) // 吉隆.阿普乔.佛特拉 克里斯汀
	return couples
}

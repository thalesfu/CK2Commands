package people

import "bufio"

var woundTraits = []string{
	"wounded",          // 受伤
	"maimed",           // 残废
	"disfigured",       // 毁容
	"mangled",          // 血肉模糊
	"one_eyed",         // 独眼
	"one_handed",       // 独手
	"one_legged",       // 独腿
	"severely_injured", // 身受重伤
	"blinded",          // 致盲
	"incapable",        // 无能
}

func cureWound(writer *bufio.Writer, peopleId int) {
	for _, trait := range woundTraits {
		removeTrait(writer, trait, peopleId)
	}
}

func CurePeopleWound(peopleId ...int) {
	var functions = []buildFunc{
		cureWound,
	}

	buildPeople("curewound", functions, peopleId...)
}

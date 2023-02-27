package people

import "bufio"

var fatherTraits = []string{
	"mastermind_theologian", // 神学的大师
	"erudite",               // 博学
	"just",                  // 正直
	"content",               // 安于现状
	"theologian",            // 神学家
	"brave",                 // 勇敢
	"zealous",               // 狂热
}

func addFatherTraits(writer *bufio.Writer, peopleId int) {
	for _, trait := range fatherTraits {
		writeAddTrait(writer, trait, peopleId)
	}
}

func addFatherLearning(writer *bufio.Writer, peopleId int) {
	writeAddLearning(writer, peopleId, 5)
	writeAddMartial(writer, peopleId, 15)
}

func getFatherFunctions() []buildFunc {
	return []buildFunc{
		removeEducationTraits,
		RemoveBadTraits,
		addCommonGoodTraits,
		addFatherTraits,
		addFatherLearning,
	}
}

func BuildFather(peopleId ...int) {
	buildPeople("father", getFatherFunctions(), peopleId...)
}

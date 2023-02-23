package people

import "bufio"

var spyTraits = []string{
	"elusive_shadow", // 难以捉摸的影子
	"cynical",        // 愤世嫉俗
	"deceitful",      // 狡诈
	"paranoid",       // 多疑
	"schemer",        // 多疑
}

func addSpyTraits(writer *bufio.Writer, peopleId int) {
	for _, trait := range spyTraits {
		addTrait(writer, trait, peopleId)
	}
}

func addSpyIntrigue(writer *bufio.Writer, peopleId int) {
	addIntrigue(writer, peopleId, 5)
}

func getSpyFunctions() []buildFunc {
	return []buildFunc{
		removeEducationTraits,
		RemoveBadTraits,
		addCommonGoodTraits,
		addSpyTraits,
		addSpyIntrigue,
	}
}

func BuildSpy(peopleId ...int) {
	buildPeople("spy", getSpyFunctions(), peopleId...)
}

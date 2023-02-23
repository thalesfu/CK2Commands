package people

import "bufio"

var generalTraits = []string{
	"brilliant_strategist", // 天才战略家
	"brave",                // 勇敢
	"zealous",              // 狂热
	"content",              // 安于现状
	"duelist",              // 决斗大师
}

func addGeneralTraits(writer *bufio.Writer, peopleId int) {
	for _, trait := range generalTraits {
		addTrait(writer, trait, peopleId)
	}
}

func addGeneralMartial(writer *bufio.Writer, peopleId int) {
	addMartial(writer, peopleId, 5)
}

func getGeneralFunctions() []buildFunc {
	return []buildFunc{
		removeEducationTraits,
		RemoveBadTraits,
		addCommonGoodTraits,
		addGeneralTraits,
		addGeneralMartial,
	}
}

func BuildGeneral(peopleId ...int) {
	buildPeople("general", getGeneralFunctions(), peopleId...)
}

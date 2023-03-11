package people

import "bufio"

var managerTraits = []string{
	"midas_touched", // 点石成金者
	"just",          // 正直
	"content",       // 安于现状
	"administrator", // 管理者
}

func addManagerTraits(writer *bufio.Writer, peopleId int) {
	for _, trait := range managerTraits {
		writeAddTrait(writer, trait, peopleId)
	}
}

func addManagerStewardship(writer *bufio.Writer, peopleId int) {
	writeAddStewardship(writer, peopleId, 5)
}

func getManagerFunctions() []buildFunc {
	return []buildFunc{
		removeEducationTraits,
		RemoveBadTraits,
		addCommonGoodTraits,
		addManagerTraits,
		addManagerStewardship,
	}
}

func BuildManager(peopleId ...int) {
	buildPeople("manager", getManagerFunctions(), peopleId...)
}

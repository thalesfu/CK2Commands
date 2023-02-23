package people

import "bufio"

var myLoadTraits = []string{
	"ambitious",     // 野心勃勃
	"brave",         // 勇敢
	"zealous",       // 狂热
	"physician",     // 训练有素的医生
	"erudite",       // 博学
	"gregarious",    // 合群
	"just",          // 正直
	"midas_touched", // 点石成金者
}

func addMyLoadTraits(writer *bufio.Writer, peopleId int) {
	for _, trait := range myLoadTraits {
		addTrait(writer, trait, peopleId)
	}
}

func addMyLoadStatus(writer *bufio.Writer, peopleId int) {
	addDiplomacy(writer, peopleId, 10)
	addMartial(writer, peopleId, 10)
	addStewardship(writer, peopleId, 10)
	addIntrigue(writer, peopleId, 10)
	addLearning(writer, peopleId, 10)
}

func getMyloadFunctions() []buildFunc {
	return []buildFunc{
		RemoveBadTraits,
		removeEducationTraits,
		addCommonGoodTraits,
		addMyLoadTraits,
		addMyLoadStatus,
	}
}

func BuildMyLoad(peopleId ...int) {
	buildPeople("myload", getMyloadFunctions(), peopleId...)
}

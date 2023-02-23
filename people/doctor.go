package people

import "bufio"

var doctorTraits = []string{
	"mastermind_theologian", // 神学的大师
	"physician",             // 训练有素的医生
	"erudite",               // 博学
	"just",                  // 正直
	"content",               // 安于现状
	"scholar",               // 学者
}

func addDoctorTraits(writer *bufio.Writer, peopleId int) {
	for _, trait := range doctorTraits {
		addTrait(writer, trait, peopleId)
	}
}

func addDoctorLearning(writer *bufio.Writer, peopleId int) {
	addLearning(writer, peopleId, 10)
}

func getDoctorFunctions() []buildFunc {
	return []buildFunc{
		removeEducationTraits,
		RemoveBadTraits,
		addCommonGoodTraits,
		addDoctorTraits,
		addDoctorLearning,
	}
}

func BuildDoctor(peopleId ...int) {
	buildPeople("doctor", getDoctorFunctions(), peopleId...)
}

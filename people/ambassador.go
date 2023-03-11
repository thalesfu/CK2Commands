package people

import (
	"bufio"
)

var ambassadorTraits = []string{
	"grey_eminence", // 幕后操控人
	"gregarious",    // 合群
	"honest",        // 诚实
	"content",       // 安于现状
	"socializer  ",  // 社交家
}

func addAmbassadorTraitsTraits(writer *bufio.Writer, peopleId int) {
	for _, trait := range ambassadorTraits {
		writeAddTrait(writer, trait, peopleId)
	}
}

func addAmbassadorDiplomacy(writer *bufio.Writer, peopleId int) {
	writeAddDiplomacy(writer, peopleId, 5)
}

func getAmbassadorFunctions() []buildFunc {
	return []buildFunc{
		removeEducationTraits,
		RemoveBadTraits,
		addCommonGoodTraits,
		addAmbassadorTraitsTraits,
		addAmbassadorDiplomacy,
	}
}

func BuildAmbassador(peopleId ...int) {
	buildPeople("amb", getAmbassadorFunctions(), peopleId...)
}

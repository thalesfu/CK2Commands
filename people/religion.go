package people

import "bufio"

func setReligionToTaoist(writer *bufio.Writer, peopleId int) {
	writeChangeReligion(writer, peopleId, "taoist")
}

func ReligionIsTaoist(peopleId ...int) {
	var functions = []buildFunc{
		setReligionToTaoist,
	}

	buildPeople("taoist", functions, peopleId...)
}

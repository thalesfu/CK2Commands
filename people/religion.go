package people

import "bufio"

func setReligionToTaoist(writer *bufio.Writer, peopleId int) {
	changeReligion(writer, peopleId, "taoist", "pictish")
}

func ReligionIsTaoist(peopleId ...int) {
	var functions = []buildFunc{
		setReligionToTaoist,
	}

	buildPeople("taoist", functions, peopleId...)
}

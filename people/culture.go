package people

import "bufio"

func setCultureToHanPictish(writer *bufio.Writer, peopleId int) {
	writeChangeCulture(writer, peopleId, "han", "pictish")
}

func CultureIsHanPictish(peopleId ...int) {
	var functions = []buildFunc{
		setCultureToHanPictish,
	}

	buildPeople("hanpictish", functions, peopleId...)
}

func setCultureToHanScottish(writer *bufio.Writer, peopleId int) {
	writeChangeCulture(writer, peopleId, "han", "scottish")
}

func CultureIsHanScottish(peopleId ...int) {
	var functions = []buildFunc{
		setCultureToHanScottish,
	}

	buildPeople("hanscottish", functions, peopleId...)
}

func setCultureToHanWelsh(writer *bufio.Writer, peopleId int) {
	writeChangeCulture(writer, peopleId, "han", "welsh")
}

func CultureIsHanWelsh(peopleId ...int) {
	var functions = []buildFunc{
		setCultureToHanWelsh,
	}

	buildPeople("hanwelsh", functions, peopleId...)
}

package people

import (
	"bufio"
	"fmt"
)

func addTrait(writer *bufio.Writer, trait string, peopleId int) {
	writer.WriteString(fmt.Sprintf("add_trait %s %d\n", trait, peopleId))
}

func removeTrait(writer *bufio.Writer, trait string, peopleId int) {
	writer.WriteString(fmt.Sprintf("remove_trait %s %d\n", trait, peopleId))
}

func addDiplomacy(writer *bufio.Writer, peopleId int, value int) {
	writer.WriteString(fmt.Sprintf("add_diplomacy %d %d\n", peopleId, value))
}

func addMartial(writer *bufio.Writer, peopleId int, value int) {
	writer.WriteString(fmt.Sprintf("add_martial %d %d\n", peopleId, value))
}

func addStewardship(writer *bufio.Writer, peopleId int, value int) {
	writer.WriteString(fmt.Sprintf("add_stewardship %d %d\n", peopleId, value))
}

func addIntrigue(writer *bufio.Writer, peopleId int, value int) {
	writer.WriteString(fmt.Sprintf("add_intrigue %d %d\n", peopleId, value))
}

func addLearning(writer *bufio.Writer, peopleId int, value int) {
	writer.WriteString(fmt.Sprintf("add_learning %d %d\n", peopleId, value))
}

func changeCulture(writer *bufio.Writer, peopleId int, culture string, ethnicity string) {
	if culture != "" {
		writer.WriteString(fmt.Sprintf("culture %d %s\n", peopleId, culture))
	}

	if ethnicity != "" {
		writer.WriteString(fmt.Sprintf("gfx_culture %d %s\n", peopleId, ethnicity))
	}
}

func changeReligion(writer *bufio.Writer, peopleId int, religion string, ethnicity string) {
	writer.WriteString(fmt.Sprintf("religion %d %s\n", peopleId, religion))
}

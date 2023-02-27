package people

import (
	"bufio"
	"fmt"
)

func writeAddTrait(writer *bufio.Writer, trait string, peopleId int) {
	writer.WriteString(fmt.Sprintf("add_trait %s %d\n", trait, peopleId))
}

func writeRemoveTrait(writer *bufio.Writer, trait string, peopleId int) {
	writer.WriteString(fmt.Sprintf("remove_trait %s %d\n", trait, peopleId))
}

func writeAddDiplomacy(writer *bufio.Writer, peopleId int, value int) {
	writer.WriteString(fmt.Sprintf("add_diplomacy %d %d\n", peopleId, value))
}

func writeAddMartial(writer *bufio.Writer, peopleId int, value int) {
	writer.WriteString(fmt.Sprintf("add_martial %d %d\n", peopleId, value))
}

func writeAddStewardship(writer *bufio.Writer, peopleId int, value int) {
	writer.WriteString(fmt.Sprintf("add_stewardship %d %d\n", peopleId, value))
}

func writeAddIntrigue(writer *bufio.Writer, peopleId int, value int) {
	writer.WriteString(fmt.Sprintf("add_intrigue %d %d\n", peopleId, value))
}

func writeAddLearning(writer *bufio.Writer, peopleId int, value int) {
	writer.WriteString(fmt.Sprintf("add_learning %d %d\n", peopleId, value))
}

func writeChangeCulture(writer *bufio.Writer, peopleId int, culture string, ethnicity string) {
	if culture != "" {
		writer.WriteString(fmt.Sprintf("culture %d %s\n", peopleId, culture))
	}

	if ethnicity != "" {
		writer.WriteString(fmt.Sprintf("gfx_culture %d %s\n", peopleId, ethnicity))
	}
}

func writeChangeReligion(writer *bufio.Writer, peopleId int, religion string) {
	writer.WriteString(fmt.Sprintf("religion %d %s\n", peopleId, religion))
}

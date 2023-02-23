package people

import (
	"CK2Commands/utils"
	"bufio"
	"fmt"
	"log"
)

type Couple struct {
	Husband int
	Wife    int
}

func Pollinate(couples []Couple, fileName string) {
	file, err := utils.OpenFile(fileName)

	if err != nil {
		log.Println(err)
		return
	}

	defer utils.CloseFile(file)

	writer := bufio.NewWriter(file)

	for i, couple := range couples {
		if i != 0 {
			writer.WriteString("\n")
		}

		writer.WriteString(fmt.Sprintf("pollinate %d %d", couple.Wife, couple.Husband))
	}

	writer.Flush()
}

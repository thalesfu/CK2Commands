package CK2Commands

import (
	"bufio"
	"github.com/thalesfu/CK2Commands/utils"
	"log"
)

func BuildScript(fileName string, generators ...ScriptGenerator) {
	file, err := utils.OpenFile(fileName)

	if err != nil {
		log.Println(err)
		return
	}

	defer utils.CloseFile(file)

	writer := bufio.NewWriter(file)

	defer func() {
		err := writer.Flush()
		if err != nil {
			log.Println(err)
		}
	}()

	for _, generator := range generators {
		scripts := generator.Generate()
		if len(scripts) == 0 {
			continue
		}

		for _, script := range scripts {
			_, err := writer.WriteString(script + "\n")
			if err != nil {
				log.Println(err)
				return
			}
		}
	}

	writer.WriteString("sound_effect = secret_cults_gain_sympathy_01")
}

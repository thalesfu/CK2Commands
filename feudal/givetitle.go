package feudal

import (
	"bufio"
	"github.com/thalesfu/CK2Commands/utils"
	"github.com/thalesfu/paradoxtools/CK2/feud"
	"log"
)

type Title struct {
	Feuds  []feud.Feud
	Holder int
}

func GiveTitle(titles []*Title) {
	file, err := utils.OpenFile("title")

	if err != nil {
		log.Println(err)
		return
	}

	defer utils.CloseFile(file)

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for _, title := range titles {
		for _, feud := range title.Feuds {
			writeGiveTitle(writer, title.Holder, feud)
		}
	}
}

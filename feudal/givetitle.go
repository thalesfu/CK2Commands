package feudal

import (
	"bufio"
	"github.com/thalesfu/CK2Commands/earth"
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

func BuildTitle() {
	var titles []*Title
	titles = append(titles, &Title{
		Holder: 2609830,
		Feuds: []feud.Feud{
			earth.Turkestan图兰.KKhiva河中().DKhiva希瓦().CGurganj玉龙杰赤(),
			earth.Turkestan图兰.KKhiva河中().DKhiva希瓦().CGurganj玉龙杰赤().BUrgench玉龙杰赤(),
			earth.Turkestan图兰.KKhiva河中().DKhiva希瓦().CGurganj玉龙杰赤().BNukus努库斯(),
		},
	})
	GiveTitle(titles)
}

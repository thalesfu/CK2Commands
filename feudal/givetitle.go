package feudal

import (
	"bufio"
	"github.com/thalesfu/CK2Commands/earth"
	"github.com/thalesfu/CK2Commands/family/fu"
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
		Holder: fu.Me,
		Feuds: []feud.Feud{
			earth.Wendish_empire文德帝国.KPomerania波美拉尼亚().DLausitz劳西茨().CAnhalt采尔布斯特(),
		},
	})
	titles = append(titles, &Title{
		Holder: 2850972,
		Feuds: []feud.Feud{
			earth.Wendish_empire文德帝国.KPomerania波美拉尼亚().DLausitz劳西茨().CAnhalt采尔布斯特(),
		},
	})
	GiveTitle(titles)
}

package cakrakuta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 建艺罗KankerBarony struct {
	feud.BaseBarony
}

var BKanker建艺罗 feud.Barony = &建艺罗KankerBarony{}

func init() {
    f := BKanker建艺罗.(*建艺罗KankerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kanker",
		TitleName: "建艺罗",
		TitleCode: "b_kanker",
	}
}

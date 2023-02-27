package bure

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 康康KankanBarony struct {
	feud.BaseBarony
}

var BKankan康康 feud.Barony = &康康KankanBarony{}

func init() {
    f := BKankan康康.(*康康KankanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kankan",
		TitleName: "康康",
		TitleCode: "b_kankan",
	}
}

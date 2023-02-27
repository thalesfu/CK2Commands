package chagai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坎达尔KandarBarony struct {
	feud.BaseBarony
}

var BKandar坎达尔 feud.Barony = &坎达尔KandarBarony{}

func init() {
    f := BKandar坎达尔.(*坎达尔KandarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kandar",
		TitleName: "坎达尔",
		TitleCode: "b_kandar",
	}
}

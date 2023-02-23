package somnath

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 甘祇陀KanjetarBarony struct {
	feud.BaseBarony
}

var BKanjetar甘祇陀 feud.Barony = &甘祇陀KanjetarBarony{}

func init() {
	f := BKanjetar甘祇陀.(*甘祇陀KanjetarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kanjetar",
		TitleName: "甘祇陀",
		TitleCode: "b_kanjetar",
	}
}

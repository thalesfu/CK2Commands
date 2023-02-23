package kundina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 军稚那KundinaBarony struct {
	feud.BaseBarony
}

var BKundina军稚那 feud.Barony = &军稚那KundinaBarony{}

func init() {
	f := BKundina军稚那.(*军稚那KundinaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kundina",
		TitleName: "军稚那",
		TitleCode: "b_kundina",
	}
}

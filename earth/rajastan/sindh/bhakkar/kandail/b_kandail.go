package kandail

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 甘代尔KandailBarony struct {
	feud.BaseBarony
}

var BKandail甘代尔 feud.Barony = &甘代尔KandailBarony{}

func init() {
    f := BKandail甘代尔.(*甘代尔KandailBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kandail",
		TitleName: "甘代尔",
		TitleCode: "b_kandail",
	}
}

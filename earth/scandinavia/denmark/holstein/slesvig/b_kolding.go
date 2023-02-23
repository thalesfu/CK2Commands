package slesvig

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科灵KoldingBarony struct {
	feud.BaseBarony
}

var BKolding科灵 feud.Barony = &科灵KoldingBarony{}

func init() {
	f := BKolding科灵.(*科灵KoldingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kolding",
		TitleName: "科灵",
		TitleCode: "b_kolding",
	}
}

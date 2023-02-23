package esztergom

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科马罗姆KomaromBarony struct {
	feud.BaseBarony
}

var BKomarom科马罗姆 feud.Barony = &科马罗姆KomaromBarony{}

func init() {
	f := BKomarom科马罗姆.(*科马罗姆KomaromBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "komarom",
		TitleName: "科马罗姆",
		TitleCode: "b_komarom",
	}
}

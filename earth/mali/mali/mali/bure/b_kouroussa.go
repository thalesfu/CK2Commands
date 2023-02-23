package bure

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库鲁萨KouroussaBarony struct {
	feud.BaseBarony
}

var BKouroussa库鲁萨 feud.Barony = &库鲁萨KouroussaBarony{}

func init() {
	f := BKouroussa库鲁萨.(*库鲁萨KouroussaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kouroussa",
		TitleName: "库鲁萨",
		TitleCode: "b_kouroussa",
	}
}

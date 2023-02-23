package karmanta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 羯曼多KarmantaBarony struct {
	feud.BaseBarony
}

var BKarmanta羯曼多 feud.Barony = &羯曼多KarmantaBarony{}

func init() {
	f := BKarmanta羯曼多.(*羯曼多KarmantaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karmanta",
		TitleName: "羯曼多",
		TitleCode: "b_karmanta",
	}
}

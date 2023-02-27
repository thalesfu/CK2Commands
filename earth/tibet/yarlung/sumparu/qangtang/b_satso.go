package qangtang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨错SatsoBarony struct {
	feud.BaseBarony
}

var BSatso萨错 feud.Barony = &萨错SatsoBarony{}

func init() {
    f := BSatso萨错.(*萨错SatsoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "satso",
		TitleName: "萨错",
		TitleCode: "b_satso",
	}
}

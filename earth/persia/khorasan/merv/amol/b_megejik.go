package amol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅盖吉MegejikBarony struct {
	feud.BaseBarony
}

var BMegejik梅盖吉 feud.Barony = &梅盖吉MegejikBarony{}

func init() {
    f := BMegejik梅盖吉.(*梅盖吉MegejikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "megejik",
		TitleName: "梅盖吉",
		TitleCode: "b_megejik",
	}
}

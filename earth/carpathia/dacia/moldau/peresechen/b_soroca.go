package peresechen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索罗卡SorocaBarony struct {
	feud.BaseBarony
}

var BSoroca索罗卡 feud.Barony = &索罗卡SorocaBarony{}

func init() {
    f := BSoroca索罗卡.(*索罗卡SorocaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "soroca",
		TitleName: "索罗卡",
		TitleCode: "b_soroca",
	}
}

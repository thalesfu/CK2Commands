package hlynov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊日卡尔IzkarBarony struct {
	feud.BaseBarony
}

var BIzkar伊日卡尔 feud.Barony = &伊日卡尔IzkarBarony{}

func init() {
    f := BIzkar伊日卡尔.(*伊日卡尔IzkarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "izkar",
		TitleName: "伊日卡尔",
		TitleCode: "b_izkar",
	}
}

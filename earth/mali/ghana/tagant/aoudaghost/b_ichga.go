package aoudaghost

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊希加IchgaBarony struct {
	feud.BaseBarony
}

var BIchga伊希加 feud.Barony = &伊希加IchgaBarony{}

func init() {
	f := BIchga伊希加.(*伊希加IchgaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ichga",
		TitleName: "伊希加",
		TitleCode: "b_ichga",
	}
}

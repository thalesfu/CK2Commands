package ui_fiachrach

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊尼什盖InishkeaBarony struct {
	feud.BaseBarony
}

var BInishkea伊尼什盖 feud.Barony = &伊尼什盖InishkeaBarony{}

func init() {
    f := BInishkea伊尼什盖.(*伊尼什盖InishkeaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "inishkea",
		TitleName: "伊尼什盖",
		TitleCode: "b_inishkea",
	}
}

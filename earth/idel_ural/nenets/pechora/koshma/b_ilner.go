package koshma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊尔涅尔IlnerBarony struct {
	feud.BaseBarony
}

var BIlner伊尔涅尔 feud.Barony = &伊尔涅尔IlnerBarony{}

func init() {
    f := BIlner伊尔涅尔.(*伊尔涅尔IlnerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ilner",
		TitleName: "伊尔涅尔",
		TitleCode: "b_ilner",
	}
}

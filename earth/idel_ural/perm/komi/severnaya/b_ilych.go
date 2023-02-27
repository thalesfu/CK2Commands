package severnaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊雷奇IlychBarony struct {
	feud.BaseBarony
}

var BIlych伊雷奇 feud.Barony = &伊雷奇IlychBarony{}

func init() {
    f := BIlych伊雷奇.(*伊雷奇IlychBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ilych",
		TitleName: "伊雷奇",
		TitleCode: "b_ilych",
	}
}

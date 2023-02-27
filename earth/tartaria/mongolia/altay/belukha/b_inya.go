package belukha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊尼亚InyaBarony struct {
	feud.BaseBarony
}

var BInya伊尼亚 feud.Barony = &伊尼亚InyaBarony{}

func init() {
    f := BInya伊尼亚.(*伊尼亚InyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "inya",
		TitleName: "伊尼亚",
		TitleCode: "b_inya",
	}
}

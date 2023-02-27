package ishim

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊希姆IshimBarony struct {
	feud.BaseBarony
}

var BIshim伊希姆 feud.Barony = &伊希姆IshimBarony{}

func init() {
    f := BIshim伊希姆.(*伊希姆IshimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ishim",
		TitleName: "伊希姆",
		TitleCode: "b_ishim",
	}
}

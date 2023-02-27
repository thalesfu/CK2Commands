package irtek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊尔捷克IrtekBarony struct {
	feud.BaseBarony
}

var BIrtek伊尔捷克 feud.Barony = &伊尔捷克IrtekBarony{}

func init() {
    f := BIrtek伊尔捷克.(*伊尔捷克IrtekBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "irtek",
		TitleName: "伊尔捷克",
		TitleCode: "b_irtek",
	}
}

package lyon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊里尼IrignyBarony struct {
	feud.BaseBarony
}

var BIrigny伊里尼 feud.Barony = &伊里尼IrignyBarony{}

func init() {
    f := BIrigny伊里尼.(*伊里尼IrignyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "irigny",
		TitleName: "伊里尼",
		TitleCode: "b_irigny",
	}
}

package aland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨尔特维克SaltvikBarony struct {
	feud.BaseBarony
}

var BSaltvik萨尔特维克 feud.Barony = &萨尔特维克SaltvikBarony{}

func init() {
    f := BSaltvik萨尔特维克.(*萨尔特维克SaltvikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saltvik",
		TitleName: "萨尔特维克",
		TitleCode: "b_saltvik",
	}
}

package saravan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨尔巴兹SarbazBarony struct {
	feud.BaseBarony
}

var BSarbaz萨尔巴兹 feud.Barony = &萨尔巴兹SarbazBarony{}

func init() {
    f := BSarbaz萨尔巴兹.(*萨尔巴兹SarbazBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarbaz",
		TitleName: "萨尔巴兹",
		TitleCode: "b_sarbaz",
	}
}

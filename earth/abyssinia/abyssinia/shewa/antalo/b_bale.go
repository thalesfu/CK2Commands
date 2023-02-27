package antalo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴莱BaleBarony struct {
	feud.BaseBarony
}

var BBale巴莱 feud.Barony = &巴莱BaleBarony{}

func init() {
    f := BBale巴莱.(*巴莱BaleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bale",
		TitleName: "巴莱",
		TitleCode: "b_bale",
	}
}

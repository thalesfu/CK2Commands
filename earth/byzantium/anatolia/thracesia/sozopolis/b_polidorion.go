package sozopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波利德良PolidorionBarony struct {
	feud.BaseBarony
}

var BPolidorion波利德良 feud.Barony = &波利德良PolidorionBarony{}

func init() {
	f := BPolidorion波利德良.(*波利德良PolidorionBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "polidorion",
		TitleName: "波利德良",
		TitleCode: "b_polidorion",
	}
}

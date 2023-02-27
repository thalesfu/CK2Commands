package onogost

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 久尔杰维斯图波维Durdevi_stupoviBarony struct {
	feud.BaseBarony
}

var BDurdevi_stupovi久尔杰维斯图波维 feud.Barony = &久尔杰维斯图波维Durdevi_stupoviBarony{}

func init() {
    f := BDurdevi_stupovi久尔杰维斯图波维.(*久尔杰维斯图波维Durdevi_stupoviBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "durdevi_stupovi",
		TitleName: "久尔杰维斯图波维",
		TitleCode: "b_durdevi_stupovi",
	}
}

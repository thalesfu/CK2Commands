package antiocheia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达巴萨克DarbsakBarony struct {
	feud.BaseBarony
}

var BDarbsak达巴萨克 feud.Barony = &达巴萨克DarbsakBarony{}

func init() {
    f := BDarbsak达巴萨克.(*达巴萨克DarbsakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "darbsak",
		TitleName: "达巴萨克",
		TitleCode: "b_darbsak",
	}
}

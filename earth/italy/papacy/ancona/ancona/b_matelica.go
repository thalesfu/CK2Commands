package ancona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马泰利卡MatelicaBarony struct {
	feud.BaseBarony
}

var BMatelica马泰利卡 feud.Barony = &马泰利卡MatelicaBarony{}

func init() {
    f := BMatelica马泰利卡.(*马泰利卡MatelicaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "matelica",
		TitleName: "马泰利卡",
		TitleCode: "b_matelica",
	}
}

package la_marche

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布萨克BoussacBarony struct {
	feud.BaseBarony
}

var BBoussac布萨克 feud.Barony = &布萨克BoussacBarony{}

func init() {
    f := BBoussac布萨克.(*布萨克BoussacBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "boussac",
		TitleName: "布萨克",
		TitleCode: "b_boussac",
	}
}

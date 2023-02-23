package kaisereia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米斯蒂MistiBarony struct {
	feud.BaseBarony
}

var BMisti米斯蒂 feud.Barony = &米斯蒂MistiBarony{}

func init() {
	f := BMisti米斯蒂.(*米斯蒂MistiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "misti",
		TitleName: "米斯蒂",
		TitleCode: "b_misti",
	}
}

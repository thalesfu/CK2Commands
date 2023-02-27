package al_djazair

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提帕萨TipasaBarony struct {
	feud.BaseBarony
}

var BTipasa提帕萨 feud.Barony = &提帕萨TipasaBarony{}

func init() {
    f := BTipasa提帕萨.(*提帕萨TipasaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tipasa",
		TitleName: "提帕萨",
		TitleCode: "b_tipasa",
	}
}

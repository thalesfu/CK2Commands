package thuringen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨尔聪根SalzungenBarony struct {
	feud.BaseBarony
}

var BSalzungen萨尔聪根 feud.Barony = &萨尔聪根SalzungenBarony{}

func init() {
    f := BSalzungen萨尔聪根.(*萨尔聪根SalzungenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "salzungen",
		TitleName: "萨尔聪根",
		TitleCode: "b_salzungen",
	}
}

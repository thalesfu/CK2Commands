package bourges

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 代奥勒DeolsBarony struct {
	feud.BaseBarony
}

var BDeols代奥勒 feud.Barony = &代奥勒DeolsBarony{}

func init() {
    f := BDeols代奥勒.(*代奥勒DeolsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "deols",
		TitleName: "代奥勒",
		TitleCode: "b_deols",
	}
}

package tadmekka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾萨科EssakoBarony struct {
	feud.BaseBarony
}

var BEssako艾萨科 feud.Barony = &艾萨科EssakoBarony{}

func init() {
    f := BEssako艾萨科.(*艾萨科EssakoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "essako",
		TitleName: "艾萨科",
		TitleCode: "b_essako",
	}
}

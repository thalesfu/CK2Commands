package crimea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨克SaqBarony struct {
	feud.BaseBarony
}

var BSaq萨克 feud.Barony = &萨克SaqBarony{}

func init() {
    f := BSaq萨克.(*萨克SaqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saq",
		TitleName: "萨克",
		TitleCode: "b_saq",
	}
}

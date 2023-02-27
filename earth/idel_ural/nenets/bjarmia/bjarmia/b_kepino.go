package bjarmia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克皮诺KepinoBarony struct {
	feud.BaseBarony
}

var BKepino克皮诺 feud.Barony = &克皮诺KepinoBarony{}

func init() {
    f := BKepino克皮诺.(*克皮诺KepinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kepino",
		TitleName: "克皮诺",
		TitleCode: "b_kepino",
	}
}

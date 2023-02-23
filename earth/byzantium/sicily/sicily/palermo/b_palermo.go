package palermo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴勒莫PalermoBarony struct {
	feud.BaseBarony
}

var BPalermo巴勒莫 feud.Barony = &巴勒莫PalermoBarony{}

func init() {
	f := BPalermo巴勒莫.(*巴勒莫PalermoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "palermo",
		TitleName: "巴勒莫",
		TitleCode: "b_palermo",
	}
}

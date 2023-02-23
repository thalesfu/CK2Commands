package ghana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪亚拉DiaraBarony struct {
	feud.BaseBarony
}

var BDiara迪亚拉 feud.Barony = &迪亚拉DiaraBarony{}

func init() {
	f := BDiara迪亚拉.(*迪亚拉DiaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "diara",
		TitleName: "迪亚拉",
		TitleCode: "b_diara",
	}
}

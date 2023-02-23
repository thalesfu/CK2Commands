package schwyz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪森蒂斯DisentisBarony struct {
	feud.BaseBarony
}

var BDisentis迪森蒂斯 feud.Barony = &迪森蒂斯DisentisBarony{}

func init() {
	f := BDisentis迪森蒂斯.(*迪森蒂斯DisentisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "disentis",
		TitleName: "迪森蒂斯",
		TitleCode: "b_disentis",
	}
}

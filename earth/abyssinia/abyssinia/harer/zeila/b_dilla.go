package zeila

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪拉DillaBarony struct {
	feud.BaseBarony
}

var BDilla迪拉 feud.Barony = &迪拉DillaBarony{}

func init() {
	f := BDilla迪拉.(*迪拉DillaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dilla",
		TitleName: "迪拉",
		TitleCode: "b_dilla",
	}
}

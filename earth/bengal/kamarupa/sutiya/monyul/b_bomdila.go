package monyul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 邦迪拉BomdilaBarony struct {
	feud.BaseBarony
}

var BBomdila邦迪拉 feud.Barony = &邦迪拉BomdilaBarony{}

func init() {
	f := BBomdila邦迪拉.(*邦迪拉BomdilaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bomdila",
		TitleName: "邦迪拉",
		TitleCode: "b_bomdila",
	}
}

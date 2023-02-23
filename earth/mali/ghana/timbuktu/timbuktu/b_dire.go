package timbuktu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪雷DireBarony struct {
	feud.BaseBarony
}

var BDire迪雷 feud.Barony = &迪雷DireBarony{}

func init() {
	f := BDire迪雷.(*迪雷DireBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dire",
		TitleName: "迪雷",
		TitleCode: "b_dire",
	}
}

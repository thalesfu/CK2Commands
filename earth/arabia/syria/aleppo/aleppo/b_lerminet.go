package aleppo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱尔米讷LerminetBarony struct {
	feud.BaseBarony
}

var BLerminet莱尔米讷 feud.Barony = &莱尔米讷LerminetBarony{}

func init() {
    f := BLerminet莱尔米讷.(*莱尔米讷LerminetBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lerminet",
		TitleName: "莱尔米讷",
		TitleCode: "b_lerminet",
	}
}

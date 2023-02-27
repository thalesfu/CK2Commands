package lykandos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 律坎多斯LykandosBarony struct {
	feud.BaseBarony
}

var BLykandos律坎多斯 feud.Barony = &律坎多斯LykandosBarony{}

func init() {
    f := BLykandos律坎多斯.(*律坎多斯LykandosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lykandos",
		TitleName: "律坎多斯",
		TitleCode: "b_lykandos",
	}
}

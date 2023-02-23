package slesvig

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 里伯RibeBarony struct {
	feud.BaseBarony
}

var BRibe里伯 feud.Barony = &里伯RibeBarony{}

func init() {
	f := BRibe里伯.(*里伯RibeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ribe",
		TitleName: "里伯",
		TitleCode: "b_ribe",
	}
}

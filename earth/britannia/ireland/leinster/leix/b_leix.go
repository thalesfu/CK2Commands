package leix

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱伊什LeixBarony struct {
	feud.BaseBarony
}

var BLeix莱伊什 feud.Barony = &莱伊什LeixBarony{}

func init() {
	f := BLeix莱伊什.(*莱伊什LeixBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "leix",
		TitleName: "莱伊什",
		TitleCode: "b_leix",
	}
}

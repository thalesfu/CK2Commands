package meath

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉莱里Raith_laoghaireBarony struct {
	feud.BaseBarony
}

var BRaith_laoghaire拉莱里 feud.Barony = &拉莱里Raith_laoghaireBarony{}

func init() {
    f := BRaith_laoghaire拉莱里.(*拉莱里Raith_laoghaireBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "raith_laoghaire",
		TitleName: "拉莱里",
		TitleCode: "b_raith_laoghaire",
	}
}

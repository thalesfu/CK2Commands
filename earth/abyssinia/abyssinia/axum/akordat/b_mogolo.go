package akordat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫格洛MogoloBarony struct {
	feud.BaseBarony
}

var BMogolo莫格洛 feud.Barony = &莫格洛MogoloBarony{}

func init() {
	f := BMogolo莫格洛.(*莫格洛MogoloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mogolo",
		TitleName: "莫格洛",
		TitleCode: "b_mogolo",
	}
}

package negev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪莫纳DimonaBarony struct {
	feud.BaseBarony
}

var BDimona迪莫纳 feud.Barony = &迪莫纳DimonaBarony{}

func init() {
	f := BDimona迪莫纳.(*迪莫纳DimonaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dimona",
		TitleName: "迪莫纳",
		TitleCode: "b_dimona",
	}
}

package ranikot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波罗PaatBarony struct {
	feud.BaseBarony
}

var BPaat波罗 feud.Barony = &波罗PaatBarony{}

func init() {
	f := BPaat波罗.(*波罗PaatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "paat",
		TitleName: "波罗",
		TitleCode: "b_paat",
	}
}

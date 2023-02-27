package kundina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毕哩体弥那揭罗PrithiminagarBarony struct {
	feud.BaseBarony
}

var BPrithiminagar毕哩体弥那揭罗 feud.Barony = &毕哩体弥那揭罗PrithiminagarBarony{}

func init() {
    f := BPrithiminagar毕哩体弥那揭罗.(*毕哩体弥那揭罗PrithiminagarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "prithiminagar",
		TitleName: "毕哩体弥那揭罗",
		TitleCode: "b_prithiminagar",
	}
}

package kamarupanagara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 东光城PragyotisapuraBarony struct {
	feud.BaseBarony
}

var BPragyotisapura东光城 feud.Barony = &东光城PragyotisapuraBarony{}

func init() {
	f := BPragyotisapura东光城.(*东光城PragyotisapuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pragyotisapura",
		TitleName: "东光城",
		TitleCode: "b_pragyotisapura",
	}
}

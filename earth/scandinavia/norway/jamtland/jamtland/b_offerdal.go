package jamtland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥弗达尔OfferdalBarony struct {
	feud.BaseBarony
}

var BOfferdal奥弗达尔 feud.Barony = &奥弗达尔OfferdalBarony{}

func init() {
	f := BOfferdal奥弗达尔.(*奥弗达尔OfferdalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "offerdal",
		TitleName: "奥弗达尔",
		TitleCode: "b_offerdal",
	}
}

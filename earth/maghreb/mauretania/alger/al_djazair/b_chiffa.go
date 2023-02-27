package al_djazair

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希法ChiffaBarony struct {
	feud.BaseBarony
}

var BChiffa希法 feud.Barony = &希法ChiffaBarony{}

func init() {
    f := BChiffa希法.(*希法ChiffaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chiffa",
		TitleName: "希法",
		TitleCode: "b_chiffa",
	}
}

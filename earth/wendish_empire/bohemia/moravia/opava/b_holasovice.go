package opava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍拉索维采HolasoviceBarony struct {
	feud.BaseBarony
}

var BHolasovice霍拉索维采 feud.Barony = &霍拉索维采HolasoviceBarony{}

func init() {
    f := BHolasovice霍拉索维采.(*霍拉索维采HolasoviceBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "holasovice",
		TitleName: "霍拉索维采",
		TitleCode: "b_holasovice",
	}
}

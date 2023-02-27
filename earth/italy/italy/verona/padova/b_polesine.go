package padova

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波莱西内PolesineBarony struct {
	feud.BaseBarony
}

var BPolesine波莱西内 feud.Barony = &波莱西内PolesineBarony{}

func init() {
    f := BPolesine波莱西内.(*波莱西内PolesineBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "polesine",
		TitleName: "波莱西内",
		TitleCode: "b_polesine",
	}
}

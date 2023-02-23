package pithapuram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼达佩塔MandapetaBarony struct {
	feud.BaseBarony
}

var BMandapeta曼达佩塔 feud.Barony = &曼达佩塔MandapetaBarony{}

func init() {
	f := BMandapeta曼达佩塔.(*曼达佩塔MandapetaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mandapeta",
		TitleName: "曼达佩塔",
		TitleCode: "b_mandapeta",
	}
}

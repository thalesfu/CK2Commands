package tukums

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨比莱SabileBarony struct {
	feud.BaseBarony
}

var BSabile萨比莱 feud.Barony = &萨比莱SabileBarony{}

func init() {
    f := BSabile萨比莱.(*萨比莱SabileBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sabile",
		TitleName: "萨比莱",
		TitleCode: "b_sabile",
	}
}

package novgorod

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索利齐SoltsyBarony struct {
	feud.BaseBarony
}

var BSoltsy索利齐 feud.Barony = &索利齐SoltsyBarony{}

func init() {
	f := BSoltsy索利齐.(*索利齐SoltsyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "soltsy",
		TitleName: "索利齐",
		TitleCode: "b_soltsy",
	}
}

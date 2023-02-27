package severin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德罗贝塔DrobetaBarony struct {
	feud.BaseBarony
}

var BDrobeta德罗贝塔 feud.Barony = &德罗贝塔DrobetaBarony{}

func init() {
    f := BDrobeta德罗贝塔.(*德罗贝塔DrobetaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "drobeta",
		TitleName: "德罗贝塔",
		TitleCode: "b_drobeta",
	}
}

package oriel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德罗赫达DroghedaBarony struct {
	feud.BaseBarony
}

var BDrogheda德罗赫达 feud.Barony = &德罗赫达DroghedaBarony{}

func init() {
	f := BDrogheda德罗赫达.(*德罗赫达DroghedaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "drogheda",
		TitleName: "德罗赫达",
		TitleCode: "b_drogheda",
	}
}

package al_djazair

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 里赫达多RhedadouaBarony struct {
	feud.BaseBarony
}

var BRhedadoua里赫达多 feud.Barony = &里赫达多RhedadouaBarony{}

func init() {
    f := BRhedadoua里赫达多.(*里赫达多RhedadouaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rhedadoua",
		TitleName: "里赫达多",
		TitleCode: "b_rhedadoua",
	}
}

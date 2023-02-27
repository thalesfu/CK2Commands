package rimini

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 里米尼RiminiBarony struct {
	feud.BaseBarony
}

var BRimini里米尼 feud.Barony = &里米尼RiminiBarony{}

func init() {
    f := BRimini里米尼.(*里米尼RiminiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rimini",
		TitleName: "里米尼",
		TitleCode: "b_rimini",
	}
}

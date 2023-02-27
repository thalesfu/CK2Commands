package tirol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 里德RiedBarony struct {
	feud.BaseBarony
}

var BRied里德 feud.Barony = &里德RiedBarony{}

func init() {
    f := BRied里德.(*里德RiedBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ried",
		TitleName: "里德",
		TitleCode: "b_ried",
	}
}

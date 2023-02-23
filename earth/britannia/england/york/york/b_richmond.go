package york

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 里奇蒙RichmondBarony struct {
	feud.BaseBarony
}

var BRichmond里奇蒙 feud.Barony = &里奇蒙RichmondBarony{}

func init() {
	f := BRichmond里奇蒙.(*里奇蒙RichmondBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "richmond",
		TitleName: "里奇蒙",
		TitleCode: "b_richmond",
	}
}

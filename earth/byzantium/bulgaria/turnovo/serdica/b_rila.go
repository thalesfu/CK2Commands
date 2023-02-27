package serdica

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 里拉RilaBarony struct {
	feud.BaseBarony
}

var BRila里拉 feud.Barony = &里拉RilaBarony{}

func init() {
    f := BRila里拉.(*里拉RilaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rila",
		TitleName: "里拉",
		TitleCode: "b_rila",
	}
}

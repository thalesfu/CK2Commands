package riga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 里加RigaBarony struct {
	feud.BaseBarony
}

var BRiga里加 feud.Barony = &里加RigaBarony{}

func init() {
    f := BRiga里加.(*里加RigaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "riga",
		TitleName: "里加",
		TitleCode: "b_riga",
	}
}

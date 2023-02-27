package qaidam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 里沟渠LigouchuBarony struct {
	feud.BaseBarony
}

var BLigouchu里沟渠 feud.Barony = &里沟渠LigouchuBarony{}

func init() {
    f := BLigouchu里沟渠.(*里沟渠LigouchuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ligouchu",
		TitleName: "里沟渠",
		TitleCode: "b_ligouchu",
	}
}

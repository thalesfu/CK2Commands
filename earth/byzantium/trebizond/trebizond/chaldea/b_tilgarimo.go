package chaldea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提格里莫TilgarimoBarony struct {
	feud.BaseBarony
}

var BTilgarimo提格里莫 feud.Barony = &提格里莫TilgarimoBarony{}

func init() {
    f := BTilgarimo提格里莫.(*提格里莫TilgarimoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tilgarimo",
		TitleName: "提格里莫",
		TitleCode: "b_tilgarimo",
	}
}

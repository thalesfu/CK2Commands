package dorostotum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨穆伊尔SamuilBarony struct {
	feud.BaseBarony
}

var BSamuil萨穆伊尔 feud.Barony = &萨穆伊尔SamuilBarony{}

func init() {
	f := BSamuil萨穆伊尔.(*萨穆伊尔SamuilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "samuil",
		TitleName: "萨穆伊尔",
		TitleCode: "b_samuil",
	}
}

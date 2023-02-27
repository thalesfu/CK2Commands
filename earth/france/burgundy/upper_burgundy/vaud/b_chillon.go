package vaud

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希永ChillonBarony struct {
	feud.BaseBarony
}

var BChillon希永 feud.Barony = &希永ChillonBarony{}

func init() {
    f := BChillon希永.(*希永ChillonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chillon",
		TitleName: "希永",
		TitleCode: "b_chillon",
	}
}

package sripuri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 室利补梨SripuriBarony struct {
	feud.BaseBarony
}

var BSripuri室利补梨 feud.Barony = &室利补梨SripuriBarony{}

func init() {
    f := BSripuri室利补梨.(*室利补梨SripuriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sripuri",
		TitleName: "室利补梨",
		TitleCode: "b_sripuri",
	}
}

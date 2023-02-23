package kinda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥珀比OppebyBarony struct {
	feud.BaseBarony
}

var BOppeby奥珀比 feud.Barony = &奥珀比OppebyBarony{}

func init() {
	f := BOppeby奥珀比.(*奥珀比OppebyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oppeby",
		TitleName: "奥珀比",
		TitleCode: "b_oppeby",
	}
}

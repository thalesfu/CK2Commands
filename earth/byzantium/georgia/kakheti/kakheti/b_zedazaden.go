package kakheti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泽达泽登ZedazadenBarony struct {
	feud.BaseBarony
}

var BZedazaden泽达泽登 feud.Barony = &泽达泽登ZedazadenBarony{}

func init() {
	f := BZedazaden泽达泽登.(*泽达泽登ZedazadenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zedazaden",
		TitleName: "泽达泽登",
		TitleCode: "b_zedazaden",
	}
}

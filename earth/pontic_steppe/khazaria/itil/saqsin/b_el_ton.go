package saqsin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃尔顿El_tonBarony struct {
	feud.BaseBarony
}

var BEl_ton埃尔顿 feud.Barony = &埃尔顿El_tonBarony{}

func init() {
    f := BEl_ton埃尔顿.(*埃尔顿El_tonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "el_ton",
		TitleName: "埃尔顿",
		TitleCode: "b_el_ton",
	}
}

package suf

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 欧韦奈El_aouinaBarony struct {
	feud.BaseBarony
}

var BEl_aouina欧韦奈 feud.Barony = &欧韦奈El_aouinaBarony{}

func init() {
    f := BEl_aouina欧韦奈.(*欧韦奈El_aouinaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "el_aouina",
		TitleName: "欧韦奈",
		TitleCode: "b_el_aouina",
	}
}

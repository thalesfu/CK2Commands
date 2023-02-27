package lhunzhub

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 旁多PoindoBarony struct {
	feud.BaseBarony
}

var BPoindo旁多 feud.Barony = &旁多PoindoBarony{}

func init() {
    f := BPoindo旁多.(*旁多PoindoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "poindo",
		TitleName: "旁多",
		TitleCode: "b_poindo",
	}
}

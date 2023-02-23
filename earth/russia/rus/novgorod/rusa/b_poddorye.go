package rusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波多里耶PoddoryeBarony struct {
	feud.BaseBarony
}

var BPoddorye波多里耶 feud.Barony = &波多里耶PoddoryeBarony{}

func init() {
	f := BPoddorye波多里耶.(*波多里耶PoddoryeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "poddorye",
		TitleName: "波多里耶",
		TitleCode: "b_poddorye",
	}
}

package szekelyfold

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥劳纽什塞克AranyosBarony struct {
	feud.BaseBarony
}

var BAranyos奥劳纽什塞克 feud.Barony = &奥劳纽什塞克AranyosBarony{}

func init() {
	f := BAranyos奥劳纽什塞克.(*奥劳纽什塞克AranyosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aranyos",
		TitleName: "奥劳纽什塞克",
		TitleCode: "b_aranyos",
	}
}

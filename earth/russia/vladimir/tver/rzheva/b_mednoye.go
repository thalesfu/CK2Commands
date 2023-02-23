package rzheva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅德诺耶MednoyeBarony struct {
	feud.BaseBarony
}

var BMednoye梅德诺耶 feud.Barony = &梅德诺耶MednoyeBarony{}

func init() {
	f := BMednoye梅德诺耶.(*梅德诺耶MednoyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mednoye",
		TitleName: "梅德诺耶",
		TitleCode: "b_mednoye",
	}
}

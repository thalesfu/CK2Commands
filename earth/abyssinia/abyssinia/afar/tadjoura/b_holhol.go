package tadjoura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍霍HolholBarony struct {
	feud.BaseBarony
}

var BHolhol霍霍 feud.Barony = &霍霍HolholBarony{}

func init() {
	f := BHolhol霍霍.(*霍霍HolholBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "holhol",
		TitleName: "霍霍",
		TitleCode: "b_holhol",
	}
}

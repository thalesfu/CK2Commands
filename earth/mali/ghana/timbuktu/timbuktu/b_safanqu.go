package timbuktu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨方克SafanquBarony struct {
	feud.BaseBarony
}

var BSafanqu萨方克 feud.Barony = &萨方克SafanquBarony{}

func init() {
	f := BSafanqu萨方克.(*萨方克SafanquBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "safanqu",
		TitleName: "萨方克",
		TitleCode: "b_safanqu",
	}
}

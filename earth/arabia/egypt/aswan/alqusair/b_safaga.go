package alqusair

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞法杰SafagaBarony struct {
	feud.BaseBarony
}

var BSafaga塞法杰 feud.Barony = &塞法杰SafagaBarony{}

func init() {
	f := BSafaga塞法杰.(*塞法杰SafagaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "safaga",
		TitleName: "塞法杰",
		TitleCode: "b_safaga",
	}
}

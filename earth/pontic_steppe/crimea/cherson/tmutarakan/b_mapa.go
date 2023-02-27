package tmutarakan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马帕MapaBarony struct {
	feud.BaseBarony
}

var BMapa马帕 feud.Barony = &马帕MapaBarony{}

func init() {
    f := BMapa马帕.(*马帕MapaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mapa",
		TitleName: "马帕",
		TitleCode: "b_mapa",
	}
}

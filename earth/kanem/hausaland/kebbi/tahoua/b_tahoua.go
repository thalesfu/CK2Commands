package tahoua

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔瓦TahouaBarony struct {
	feud.BaseBarony
}

var BTahoua塔瓦 feud.Barony = &塔瓦TahouaBarony{}

func init() {
	f := BTahoua塔瓦.(*塔瓦TahouaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tahoua",
		TitleName: "塔瓦",
		TitleCode: "b_tahoua",
	}
}

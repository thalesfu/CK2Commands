package gilan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔雷什TalyshBarony struct {
	feud.BaseBarony
}

var BTalysh塔雷什 feud.Barony = &塔雷什TalyshBarony{}

func init() {
	f := BTalysh塔雷什.(*塔雷什TalyshBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "talysh",
		TitleName: "塔雷什",
		TitleCode: "b_talysh",
	}
}

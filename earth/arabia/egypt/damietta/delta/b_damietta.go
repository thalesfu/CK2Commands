package delta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达米埃塔DamiettaBarony struct {
	feud.BaseBarony
}

var BDamietta达米埃塔 feud.Barony = &达米埃塔DamiettaBarony{}

func init() {
	f := BDamietta达米埃塔.(*达米埃塔DamiettaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "damietta",
		TitleName: "达米埃塔",
		TitleCode: "b_damietta",
	}
}

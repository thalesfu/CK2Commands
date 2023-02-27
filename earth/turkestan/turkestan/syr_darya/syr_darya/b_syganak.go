package syr_darya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 昔格纳黑SyganakBarony struct {
	feud.BaseBarony
}

var BSyganak昔格纳黑 feud.Barony = &昔格纳黑SyganakBarony{}

func init() {
    f := BSyganak昔格纳黑.(*昔格纳黑SyganakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "syganak",
		TitleName: "昔格纳黑",
		TitleCode: "b_syganak",
	}
}

package calatayud

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达罗卡DarocaBarony struct {
	feud.BaseBarony
}

var BDaroca达罗卡 feud.Barony = &达罗卡DarocaBarony{}

func init() {
    f := BDaroca达罗卡.(*达罗卡DarocaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "daroca",
		TitleName: "达罗卡",
		TitleCode: "b_daroca",
	}
}

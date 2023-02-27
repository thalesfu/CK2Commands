package silves

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉古什LagosBarony struct {
	feud.BaseBarony
}

var BLagos拉古什 feud.Barony = &拉古什LagosBarony{}

func init() {
    f := BLagos拉古什.(*拉古什LagosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lagos",
		TitleName: "拉古什",
		TitleCode: "b_lagos",
	}
}

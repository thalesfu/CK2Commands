package regensburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱兴费尔德LerchenfeldBarony struct {
	feud.BaseBarony
}

var BLerchenfeld莱兴费尔德 feud.Barony = &莱兴费尔德LerchenfeldBarony{}

func init() {
    f := BLerchenfeld莱兴费尔德.(*莱兴费尔德LerchenfeldBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lerchenfeld",
		TitleName: "莱兴费尔德",
		TitleCode: "b_lerchenfeld",
	}
}

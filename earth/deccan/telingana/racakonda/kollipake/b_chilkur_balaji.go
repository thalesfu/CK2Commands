package kollipake

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赤库尔巴拉吉Chilkur_balajiBarony struct {
	feud.BaseBarony
}

var BChilkur_balaji赤库尔巴拉吉 feud.Barony = &赤库尔巴拉吉Chilkur_balajiBarony{}

func init() {
    f := BChilkur_balaji赤库尔巴拉吉.(*赤库尔巴拉吉Chilkur_balajiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chilkur_balaji",
		TitleName: "赤库尔巴拉吉",
		TitleCode: "b_chilkur_balaji",
	}
}

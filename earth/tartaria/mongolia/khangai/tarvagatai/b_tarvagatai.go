package tarvagatai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔尔巴嘎泰TarvagataiBarony struct {
	feud.BaseBarony
}

var BTarvagatai塔尔巴嘎泰 feud.Barony = &塔尔巴嘎泰TarvagataiBarony{}

func init() {
	f := BTarvagatai塔尔巴嘎泰.(*塔尔巴嘎泰TarvagataiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tarvagatai",
		TitleName: "塔尔巴嘎泰",
		TitleCode: "b_tarvagatai",
	}
}

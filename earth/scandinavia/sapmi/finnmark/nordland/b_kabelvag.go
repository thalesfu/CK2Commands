package nordland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡伯尔沃格KabelvagBarony struct {
	feud.BaseBarony
}

var BKabelvag卡伯尔沃格 feud.Barony = &卡伯尔沃格KabelvagBarony{}

func init() {
	f := BKabelvag卡伯尔沃格.(*卡伯尔沃格KabelvagBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kabelvag",
		TitleName: "卡伯尔沃格",
		TitleCode: "b_kabelvag",
	}
}

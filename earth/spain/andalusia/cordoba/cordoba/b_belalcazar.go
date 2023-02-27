package cordoba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝拉尔卡萨BelalcazarBarony struct {
	feud.BaseBarony
}

var BBelalcazar贝拉尔卡萨 feud.Barony = &贝拉尔卡萨BelalcazarBarony{}

func init() {
    f := BBelalcazar贝拉尔卡萨.(*贝拉尔卡萨BelalcazarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "belalcazar",
		TitleName: "贝拉尔卡萨",
		TitleCode: "b_belalcazar",
	}
}

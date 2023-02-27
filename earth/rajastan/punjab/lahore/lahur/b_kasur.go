package lahur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡苏尔KasurBarony struct {
	feud.BaseBarony
}

var BKasur卡苏尔 feud.Barony = &卡苏尔KasurBarony{}

func init() {
    f := BKasur卡苏尔.(*卡苏尔KasurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kasur",
		TitleName: "卡苏尔",
		TitleCode: "b_kasur",
	}
}

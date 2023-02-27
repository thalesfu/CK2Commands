package kassala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡萨拉KassalaBarony struct {
	feud.BaseBarony
}

var BKassala卡萨拉 feud.Barony = &卡萨拉KassalaBarony{}

func init() {
    f := BKassala卡萨拉.(*卡萨拉KassalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kassala",
		TitleName: "卡萨拉",
		TitleCode: "b_kassala",
	}
}

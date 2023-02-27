package bologna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴尼亚卡瓦洛BagnacavalloBarony struct {
	feud.BaseBarony
}

var BBagnacavallo巴尼亚卡瓦洛 feud.Barony = &巴尼亚卡瓦洛BagnacavalloBarony{}

func init() {
    f := BBagnacavallo巴尼亚卡瓦洛.(*巴尼亚卡瓦洛BagnacavalloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bagnacavallo",
		TitleName: "巴尼亚卡瓦洛",
		TitleCode: "b_bagnacavallo",
	}
}

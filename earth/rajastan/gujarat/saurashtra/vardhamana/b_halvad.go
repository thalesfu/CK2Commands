package vardhamana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈尔瓦HalvadBarony struct {
	feud.BaseBarony
}

var BHalvad哈尔瓦 feud.Barony = &哈尔瓦HalvadBarony{}

func init() {
    f := BHalvad哈尔瓦.(*哈尔瓦HalvadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "halvad",
		TitleName: "哈尔瓦",
		TitleCode: "b_halvad",
	}
}

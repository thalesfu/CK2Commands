package zagreb

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔洛瓦茨KarlovacBarony struct {
	feud.BaseBarony
}

var BKarlovac卡尔洛瓦茨 feud.Barony = &卡尔洛瓦茨KarlovacBarony{}

func init() {
    f := BKarlovac卡尔洛瓦茨.(*卡尔洛瓦茨KarlovacBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karlovac",
		TitleName: "卡尔洛瓦茨",
		TitleCode: "b_karlovac",
	}
}

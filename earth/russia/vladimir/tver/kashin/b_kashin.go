package kashin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡申KashinBarony struct {
	feud.BaseBarony
}

var BKashin卡申 feud.Barony = &卡申KashinBarony{}

func init() {
    f := BKashin卡申.(*卡申KashinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kashin",
		TitleName: "卡申",
		TitleCode: "b_kashin",
	}
}

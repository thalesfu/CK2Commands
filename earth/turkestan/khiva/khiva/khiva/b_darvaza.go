package khiva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达瓦扎DarvazaBarony struct {
	feud.BaseBarony
}

var BDarvaza达瓦扎 feud.Barony = &达瓦扎DarvazaBarony{}

func init() {
	f := BDarvaza达瓦扎.(*达瓦扎DarvazaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "darvaza",
		TitleName: "达瓦扎",
		TitleCode: "b_darvaza",
	}
}

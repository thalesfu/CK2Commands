package deir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷瓦RawaBarony struct {
	feud.BaseBarony
}

var BRawa雷瓦 feud.Barony = &雷瓦RawaBarony{}

func init() {
    f := BRawa雷瓦.(*雷瓦RawaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rawa",
		TitleName: "雷瓦",
		TitleCode: "b_rawa",
	}
}

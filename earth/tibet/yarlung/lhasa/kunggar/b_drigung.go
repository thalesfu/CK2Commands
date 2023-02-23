package kunggar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 必力工瓦DrigungBarony struct {
	feud.BaseBarony
}

var BDrigung必力工瓦 feud.Barony = &必力工瓦DrigungBarony{}

func init() {
	f := BDrigung必力工瓦.(*必力工瓦DrigungBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "drigung",
		TitleName: "必力工瓦",
		TitleCode: "b_drigung",
	}
}

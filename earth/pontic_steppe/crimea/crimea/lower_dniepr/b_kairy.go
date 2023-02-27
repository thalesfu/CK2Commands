package lower_dniepr

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡伊里KairyBarony struct {
	feud.BaseBarony
}

var BKairy卡伊里 feud.Barony = &卡伊里KairyBarony{}

func init() {
    f := BKairy卡伊里.(*卡伊里KairyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kairy",
		TitleName: "卡伊里",
		TitleCode: "b_kairy",
	}
}

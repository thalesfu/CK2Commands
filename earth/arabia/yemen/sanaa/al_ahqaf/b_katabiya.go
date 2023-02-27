package al_ahqaf

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡塔比亚KatabiyaBarony struct {
	feud.BaseBarony
}

var BKatabiya卡塔比亚 feud.Barony = &卡塔比亚KatabiyaBarony{}

func init() {
    f := BKatabiya卡塔比亚.(*卡塔比亚KatabiyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "katabiya",
		TitleName: "卡塔比亚",
		TitleCode: "b_katabiya",
	}
}

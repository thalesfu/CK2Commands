package ghazna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡塔瓦兹KatawazBarony struct {
	feud.BaseBarony
}

var BKatawaz卡塔瓦兹 feud.Barony = &卡塔瓦兹KatawazBarony{}

func init() {
    f := BKatawaz卡塔瓦兹.(*卡塔瓦兹KatawazBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "katawaz",
		TitleName: "卡塔瓦兹",
		TitleCode: "b_katawaz",
	}
}

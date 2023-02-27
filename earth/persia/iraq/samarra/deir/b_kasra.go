package deir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡撕拉KasraBarony struct {
	feud.BaseBarony
}

var BKasra卡撕拉 feud.Barony = &卡撕拉KasraBarony{}

func init() {
    f := BKasra卡撕拉.(*卡撕拉KasraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kasra",
		TitleName: "卡撕拉",
		TitleCode: "b_kasra",
	}
}

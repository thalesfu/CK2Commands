package idatarainadu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡姆摩塔杜尔加KammatadurgaBarony struct {
	feud.BaseBarony
}

var BKammatadurga卡姆摩塔杜尔加 feud.Barony = &卡姆摩塔杜尔加KammatadurgaBarony{}

func init() {
    f := BKammatadurga卡姆摩塔杜尔加.(*卡姆摩塔杜尔加KammatadurgaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kammatadurga",
		TitleName: "卡姆摩塔杜尔加",
		TitleCode: "b_kammatadurga",
	}
}

package asas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔布卡TabqaBarony struct {
	feud.BaseBarony
}

var BTabqa塔布卡 feud.Barony = &塔布卡TabqaBarony{}

func init() {
	f := BTabqa塔布卡.(*塔布卡TabqaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tabqa",
		TitleName: "塔布卡",
		TitleCode: "b_tabqa",
	}
}

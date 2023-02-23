package calatayud

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉塔尤CalatayudBarony struct {
	feud.BaseBarony
}

var BCalatayud卡拉塔尤 feud.Barony = &卡拉塔尤CalatayudBarony{}

func init() {
	f := BCalatayud卡拉塔尤.(*卡拉塔尤CalatayudBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "calatayud",
		TitleName: "卡拉塔尤",
		TitleCode: "b_calatayud",
	}
}

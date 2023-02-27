package saur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉塔尔Karatal_saurBarony struct {
	feud.BaseBarony
}

var BKaratal_saur卡拉塔尔 feud.Barony = &卡拉塔尔Karatal_saurBarony{}

func init() {
    f := BKaratal_saur卡拉塔尔.(*卡拉塔尔Karatal_saurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karatal_saur",
		TitleName: "卡拉塔尔",
		TitleCode: "b_karatal_saur",
	}
}

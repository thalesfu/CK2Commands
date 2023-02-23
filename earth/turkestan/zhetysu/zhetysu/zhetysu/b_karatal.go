package zhetysu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉塔尔KaratalBarony struct {
	feud.BaseBarony
}

var BKaratal卡拉塔尔 feud.Barony = &卡拉塔尔KaratalBarony{}

func init() {
	f := BKaratal卡拉塔尔.(*卡拉塔尔KaratalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karatal",
		TitleName: "卡拉塔尔",
		TitleCode: "b_karatal",
	}
}

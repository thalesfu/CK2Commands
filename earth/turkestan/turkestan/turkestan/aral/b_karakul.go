package aral

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉库尔KarakulBarony struct {
	feud.BaseBarony
}

var BKarakul卡拉库尔 feud.Barony = &卡拉库尔KarakulBarony{}

func init() {
	f := BKarakul卡拉库尔.(*卡拉库尔KarakulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karakul",
		TitleName: "卡拉库尔",
		TitleCode: "b_karakul",
	}
}

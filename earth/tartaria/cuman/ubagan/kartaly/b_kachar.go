package kartaly

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡恰尔KacharBarony struct {
	feud.BaseBarony
}

var BKachar卡恰尔 feud.Barony = &卡恰尔KacharBarony{}

func init() {
    f := BKachar卡恰尔.(*卡恰尔KacharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kachar",
		TitleName: "卡恰尔",
		TitleCode: "b_kachar",
	}
}

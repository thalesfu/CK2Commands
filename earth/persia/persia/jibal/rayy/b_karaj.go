package rayy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉季KarajBarony struct {
	feud.BaseBarony
}

var BKaraj卡拉季 feud.Barony = &卡拉季KarajBarony{}

func init() {
	f := BKaraj卡拉季.(*卡拉季KarajBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karaj",
		TitleName: "卡拉季",
		TitleCode: "b_karaj",
	}
}

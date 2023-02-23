package karelen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡列瓦拉KalevalaBarony struct {
	feud.BaseBarony
}

var BKalevala卡列瓦拉 feud.Barony = &卡列瓦拉KalevalaBarony{}

func init() {
	f := BKalevala卡列瓦拉.(*卡列瓦拉KalevalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalevala",
		TitleName: "卡列瓦拉",
		TitleCode: "b_kalevala",
	}
}

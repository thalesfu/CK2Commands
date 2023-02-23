package chagai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉瓦拉KallarwalaBarony struct {
	feud.BaseBarony
}

var BKallarwala卡拉瓦拉 feud.Barony = &卡拉瓦拉KallarwalaBarony{}

func init() {
	f := BKallarwala卡拉瓦拉.(*卡拉瓦拉KallarwalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kallarwala",
		TitleName: "卡拉瓦拉",
		TitleCode: "b_kallarwala",
	}
}

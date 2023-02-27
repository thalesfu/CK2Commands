package kalaus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡劳斯KalausBarony struct {
	feud.BaseBarony
}

var BKalaus卡劳斯 feud.Barony = &卡劳斯KalausBarony{}

func init() {
    f := BKalaus卡劳斯.(*卡劳斯KalausBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalaus",
		TitleName: "卡劳斯",
		TitleCode: "b_kalaus",
	}
}

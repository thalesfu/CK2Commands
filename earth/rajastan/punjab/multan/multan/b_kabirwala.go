package multan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡比瓦拉KabirwalaBarony struct {
	feud.BaseBarony
}

var BKabirwala卡比瓦拉 feud.Barony = &卡比瓦拉KabirwalaBarony{}

func init() {
    f := BKabirwala卡比瓦拉.(*卡比瓦拉KabirwalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kabirwala",
		TitleName: "卡比瓦拉",
		TitleCode: "b_kabirwala",
	}
}

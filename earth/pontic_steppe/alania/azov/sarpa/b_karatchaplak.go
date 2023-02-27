package sarpa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉查普拉克KaratchaplakBarony struct {
	feud.BaseBarony
}

var BKaratchaplak卡拉查普拉克 feud.Barony = &卡拉查普拉克KaratchaplakBarony{}

func init() {
    f := BKaratchaplak卡拉查普拉克.(*卡拉查普拉克KaratchaplakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karatchaplak",
		TitleName: "卡拉查普拉克",
		TitleCode: "b_karatchaplak",
	}
}

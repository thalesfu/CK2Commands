package qom

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡哈克KahakBarony struct {
	feud.BaseBarony
}

var BKahak卡哈克 feud.Barony = &卡哈克KahakBarony{}

func init() {
    f := BKahak卡哈克.(*卡哈克KahakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kahak",
		TitleName: "卡哈克",
		TitleCode: "b_kahak",
	}
}

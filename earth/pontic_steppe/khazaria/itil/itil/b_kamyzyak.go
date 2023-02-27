package itil

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡梅贾克KamyzyakBarony struct {
	feud.BaseBarony
}

var BKamyzyak卡梅贾克 feud.Barony = &卡梅贾克KamyzyakBarony{}

func init() {
    f := BKamyzyak卡梅贾克.(*卡梅贾克KamyzyakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kamyzyak",
		TitleName: "卡梅贾克",
		TitleCode: "b_kamyzyak",
	}
}

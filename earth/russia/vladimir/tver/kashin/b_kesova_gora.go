package kashin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克索瓦戈拉Kesova_goraBarony struct {
	feud.BaseBarony
}

var BKesova_gora克索瓦戈拉 feud.Barony = &克索瓦戈拉Kesova_goraBarony{}

func init() {
    f := BKesova_gora克索瓦戈拉.(*克索瓦戈拉Kesova_goraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kesova_gora",
		TitleName: "克索瓦戈拉",
		TitleCode: "b_kesova_gora",
	}
}

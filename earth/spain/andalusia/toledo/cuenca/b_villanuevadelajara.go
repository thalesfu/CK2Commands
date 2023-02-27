package cuenca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈拉新镇VillanuevadelajaraBarony struct {
	feud.BaseBarony
}

var BVillanuevadelajara哈拉新镇 feud.Barony = &哈拉新镇VillanuevadelajaraBarony{}

func init() {
    f := BVillanuevadelajara哈拉新镇.(*哈拉新镇VillanuevadelajaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "villanuevadelajara",
		TitleName: "哈拉新镇",
		TitleCode: "b_villanuevadelajara",
	}
}

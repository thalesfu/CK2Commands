package coqen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达雄DawaxungBarony struct {
	feud.BaseBarony
}

var BDawaxung达雄 feud.Barony = &达雄DawaxungBarony{}

func init() {
    f := BDawaxung达雄.(*达雄DawaxungBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dawaxung",
		TitleName: "达雄",
		TitleCode: "b_dawaxung",
	}
}

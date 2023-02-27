package udmurts

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨尔坎SarkanBarony struct {
	feud.BaseBarony
}

var BSarkan萨尔坎 feud.Barony = &萨尔坎SarkanBarony{}

func init() {
    f := BSarkan萨尔坎.(*萨尔坎SarkanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarkan",
		TitleName: "萨尔坎",
		TitleCode: "b_sarkan",
	}
}

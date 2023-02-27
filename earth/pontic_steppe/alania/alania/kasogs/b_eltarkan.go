package kasogs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 厄尔塔拉克EltarkanBarony struct {
	feud.BaseBarony
}

var BEltarkan厄尔塔拉克 feud.Barony = &厄尔塔拉克EltarkanBarony{}

func init() {
    f := BEltarkan厄尔塔拉克.(*厄尔塔拉克EltarkanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eltarkan",
		TitleName: "厄尔塔拉克",
		TitleCode: "b_eltarkan",
	}
}

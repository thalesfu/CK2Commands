package french_leon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布雷莱BrelesBarony struct {
	feud.BaseBarony
}

var BBreles布雷莱 feud.Barony = &布雷莱BrelesBarony{}

func init() {
    f := BBreles布雷莱.(*布雷莱BrelesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "breles",
		TitleName: "布雷莱",
		TitleCode: "b_breles",
	}
}

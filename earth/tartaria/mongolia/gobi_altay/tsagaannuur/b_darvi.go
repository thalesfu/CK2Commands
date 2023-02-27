package tsagaannuur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达尔维DarviBarony struct {
	feud.BaseBarony
}

var BDarvi达尔维 feud.Barony = &达尔维DarviBarony{}

func init() {
    f := BDarvi达尔维.(*达尔维DarviBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "darvi",
		TitleName: "达尔维",
		TitleCode: "b_darvi",
	}
}

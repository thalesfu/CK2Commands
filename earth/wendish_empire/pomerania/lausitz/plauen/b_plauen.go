package plauen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普劳恩PlauenBarony struct {
	feud.BaseBarony
}

var BPlauen普劳恩 feud.Barony = &普劳恩PlauenBarony{}

func init() {
    f := BPlauen普劳恩.(*普劳恩PlauenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "plauen",
		TitleName: "普劳恩",
		TitleCode: "b_plauen",
	}
}

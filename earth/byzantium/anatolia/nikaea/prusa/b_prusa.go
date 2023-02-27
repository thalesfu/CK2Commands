package prusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普鲁萨PrusaBarony struct {
	feud.BaseBarony
}

var BPrusa普鲁萨 feud.Barony = &普鲁萨PrusaBarony{}

func init() {
    f := BPrusa普鲁萨.(*普鲁萨PrusaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "prusa",
		TitleName: "普鲁萨",
		TitleCode: "b_prusa",
	}
}

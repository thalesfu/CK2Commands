package nordlendingafjordungur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 胡萨维克HusavikBarony struct {
	feud.BaseBarony
}

var BHusavik胡萨维克 feud.Barony = &胡萨维克HusavikBarony{}

func init() {
    f := BHusavik胡萨维克.(*胡萨维克HusavikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "husavik",
		TitleName: "胡萨维克",
		TitleCode: "b_husavik",
	}
}

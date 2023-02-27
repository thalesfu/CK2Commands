package kyzikos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克雷马斯蒂KremastiBarony struct {
	feud.BaseBarony
}

var BKremasti克雷马斯蒂 feud.Barony = &克雷马斯蒂KremastiBarony{}

func init() {
    f := BKremasti克雷马斯蒂.(*克雷马斯蒂KremastiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kremasti",
		TitleName: "克雷马斯蒂",
		TitleCode: "b_kremasti",
	}
}

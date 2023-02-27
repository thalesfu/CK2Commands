package coruna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比利亚尔瓦VillalbaBarony struct {
	feud.BaseBarony
}

var BVillalba比利亚尔瓦 feud.Barony = &比利亚尔瓦VillalbaBarony{}

func init() {
    f := BVillalba比利亚尔瓦.(*比利亚尔瓦VillalbaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "villalba",
		TitleName: "比利亚尔瓦",
		TitleCode: "b_villalba",
	}
}

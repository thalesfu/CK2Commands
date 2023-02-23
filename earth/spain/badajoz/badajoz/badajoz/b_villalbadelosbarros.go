package badajoz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比利亚尔瓦德洛斯巴罗斯VillalbadelosbarrosBarony struct {
	feud.BaseBarony
}

var BVillalbadelosbarros比利亚尔瓦德洛斯巴罗斯 feud.Barony = &比利亚尔瓦德洛斯巴罗斯VillalbadelosbarrosBarony{}

func init() {
	f := BVillalbadelosbarros比利亚尔瓦德洛斯巴罗斯.(*比利亚尔瓦德洛斯巴罗斯VillalbadelosbarrosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "villalbadelosbarros",
		TitleName: "比利亚尔瓦德洛斯巴罗斯",
		TitleCode: "b_villalbadelosbarros",
	}
}

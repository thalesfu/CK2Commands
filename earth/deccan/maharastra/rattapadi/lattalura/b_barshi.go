package lattalura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴尔斯希BarshiBarony struct {
	feud.BaseBarony
}

var BBarshi巴尔斯希 feud.Barony = &巴尔斯希BarshiBarony{}

func init() {
    f := BBarshi巴尔斯希.(*巴尔斯希BarshiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "barshi",
		TitleName: "巴尔斯希",
		TitleCode: "b_barshi",
	}
}

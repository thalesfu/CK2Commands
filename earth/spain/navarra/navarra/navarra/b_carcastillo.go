package navarra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔卡斯蒂略CarcastilloBarony struct {
	feud.BaseBarony
}

var BCarcastillo卡尔卡斯蒂略 feud.Barony = &卡尔卡斯蒂略CarcastilloBarony{}

func init() {
    f := BCarcastillo卡尔卡斯蒂略.(*卡尔卡斯蒂略CarcastilloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "carcastillo",
		TitleName: "卡尔卡斯蒂略",
		TitleCode: "b_carcastillo",
	}
}

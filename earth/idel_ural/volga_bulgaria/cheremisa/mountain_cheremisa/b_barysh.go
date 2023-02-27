package mountain_cheremisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴雷什BaryshBarony struct {
	feud.BaseBarony
}

var BBarysh巴雷什 feud.Barony = &巴雷什BaryshBarony{}

func init() {
    f := BBarysh巴雷什.(*巴雷什BaryshBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "barysh",
		TitleName: "巴雷什",
		TitleCode: "b_barysh",
	}
}

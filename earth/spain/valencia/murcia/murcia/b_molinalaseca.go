package murcia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉塞卡MolinalasecaBarony struct {
	feud.BaseBarony
}

var BMolinalaseca拉塞卡 feud.Barony = &拉塞卡MolinalasecaBarony{}

func init() {
    f := BMolinalaseca拉塞卡.(*拉塞卡MolinalasecaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "molinalaseca",
		TitleName: "拉塞卡",
		TitleCode: "b_molinalaseca",
	}
}

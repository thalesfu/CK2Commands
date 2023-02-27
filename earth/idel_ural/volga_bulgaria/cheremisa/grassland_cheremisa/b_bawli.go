package grassland_cheremisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴夫雷BawliBarony struct {
	feud.BaseBarony
}

var BBawli巴夫雷 feud.Barony = &巴夫雷BawliBarony{}

func init() {
    f := BBawli巴夫雷.(*巴夫雷BawliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bawli",
		TitleName: "巴夫雷",
		TitleCode: "b_bawli",
	}
}

package dimapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提摩补罗DimapurBarony struct {
	feud.BaseBarony
}

var BDimapur提摩补罗 feud.Barony = &提摩补罗DimapurBarony{}

func init() {
    f := BDimapur提摩补罗.(*提摩补罗DimapurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dimapur",
		TitleName: "提摩补罗",
		TitleCode: "b_dimapur",
	}
}

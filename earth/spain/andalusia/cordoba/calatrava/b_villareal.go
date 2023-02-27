package calatrava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比利亚雷亚尔VillarealBarony struct {
	feud.BaseBarony
}

var BVillareal比利亚雷亚尔 feud.Barony = &比利亚雷亚尔VillarealBarony{}

func init() {
    f := BVillareal比利亚雷亚尔.(*比利亚雷亚尔VillarealBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "villareal",
		TitleName: "比利亚雷亚尔",
		TitleCode: "b_villareal",
	}
}

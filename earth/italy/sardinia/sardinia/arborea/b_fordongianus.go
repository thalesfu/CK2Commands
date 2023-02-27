package arborea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 福尔东贾努斯FordongianusBarony struct {
	feud.BaseBarony
}

var BFordongianus福尔东贾努斯 feud.Barony = &福尔东贾努斯FordongianusBarony{}

func init() {
    f := BFordongianus福尔东贾努斯.(*福尔东贾努斯FordongianusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fordongianus",
		TitleName: "福尔东贾努斯",
		TitleCode: "b_fordongianus",
	}
}

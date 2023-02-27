package upper_silesia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱格尼察LegnicaBarony struct {
	feud.BaseBarony
}

var BLegnica莱格尼察 feud.Barony = &莱格尼察LegnicaBarony{}

func init() {
    f := BLegnica莱格尼察.(*莱格尼察LegnicaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "legnica",
		TitleName: "莱格尼察",
		TitleCode: "b_legnica",
	}
}

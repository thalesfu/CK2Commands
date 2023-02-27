package bolshaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索伊加SoygaBarony struct {
	feud.BaseBarony
}

var BSoyga索伊加 feud.Barony = &索伊加SoygaBarony{}

func init() {
    f := BSoyga索伊加.(*索伊加SoygaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "soyga",
		TitleName: "索伊加",
		TitleCode: "b_soyga",
	}
}

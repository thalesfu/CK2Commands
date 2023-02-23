package kandail

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 威龙WilongBarony struct {
	feud.BaseBarony
}

var BWilong威龙 feud.Barony = &威龙WilongBarony{}

func init() {
	f := BWilong威龙.(*威龙WilongBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wilong",
		TitleName: "威龙",
		TitleCode: "b_wilong",
	}
}

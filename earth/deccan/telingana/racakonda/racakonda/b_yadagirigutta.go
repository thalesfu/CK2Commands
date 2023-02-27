package racakonda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶陀耆厘矩吒YadagiriguttaBarony struct {
	feud.BaseBarony
}

var BYadagirigutta耶陀耆厘矩吒 feud.Barony = &耶陀耆厘矩吒YadagiriguttaBarony{}

func init() {
    f := BYadagirigutta耶陀耆厘矩吒.(*耶陀耆厘矩吒YadagiriguttaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yadagirigutta",
		TitleName: "耶陀耆厘矩吒",
		TitleCode: "b_yadagirigutta",
	}
}

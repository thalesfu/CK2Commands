package madaba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦迪西尔Wadi_al_sirBarony struct {
	feud.BaseBarony
}

var BWadi_al_sir瓦迪西尔 feud.Barony = &瓦迪西尔Wadi_al_sirBarony{}

func init() {
    f := BWadi_al_sir瓦迪西尔.(*瓦迪西尔Wadi_al_sirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wadi_al_sir",
		TitleName: "瓦迪西尔",
		TitleCode: "b_wadi_al_sir",
	}
}

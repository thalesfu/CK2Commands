package kolguyev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦利卡亚VelikayaBarony struct {
	feud.BaseBarony
}

var BVelikaya韦利卡亚 feud.Barony = &韦利卡亚VelikayaBarony{}

func init() {
    f := BVelikaya韦利卡亚.(*韦利卡亚VelikayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "velikaya",
		TitleName: "韦利卡亚",
		TitleCode: "b_velikaya",
	}
}

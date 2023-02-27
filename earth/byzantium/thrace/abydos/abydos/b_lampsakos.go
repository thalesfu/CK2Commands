package abydos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兰姆萨科洛斯LampsakosBarony struct {
	feud.BaseBarony
}

var BLampsakos兰姆萨科洛斯 feud.Barony = &兰姆萨科洛斯LampsakosBarony{}

func init() {
    f := BLampsakos兰姆萨科洛斯.(*兰姆萨科洛斯LampsakosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lampsakos",
		TitleName: "兰姆萨科洛斯",
		TitleCode: "b_lampsakos",
	}
}

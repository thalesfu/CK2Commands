package halogaland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦加VeigaBarony struct {
	feud.BaseBarony
}

var BVeiga韦加 feud.Barony = &韦加VeigaBarony{}

func init() {
	f := BVeiga韦加.(*韦加VeigaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "veiga",
		TitleName: "韦加",
		TitleCode: "b_veiga",
	}
}

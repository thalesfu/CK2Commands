package astibus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦莱斯VelesBarony struct {
	feud.BaseBarony
}

var BVeles韦莱斯 feud.Barony = &韦莱斯VelesBarony{}

func init() {
	f := BVeles韦莱斯.(*韦莱斯VelesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "veles",
		TitleName: "韦莱斯",
		TitleCode: "b_veles",
	}
}

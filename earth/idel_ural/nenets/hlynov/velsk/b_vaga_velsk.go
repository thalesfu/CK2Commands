package velsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦加Vaga_velskBarony struct {
	feud.BaseBarony
}

var BVaga_velsk瓦加 feud.Barony = &瓦加Vaga_velskBarony{}

func init() {
    f := BVaga_velsk瓦加.(*瓦加Vaga_velskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vaga_velsk",
		TitleName: "瓦加",
		TitleCode: "b_vaga_velsk",
	}
}

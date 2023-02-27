package tavasts

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦西拉赫蒂VesilahtiBarony struct {
	feud.BaseBarony
}

var BVesilahti韦西拉赫蒂 feud.Barony = &韦西拉赫蒂VesilahtiBarony{}

func init() {
    f := BVesilahti韦西拉赫蒂.(*韦西拉赫蒂VesilahtiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vesilahti",
		TitleName: "韦西拉赫蒂",
		TitleCode: "b_vesilahti",
	}
}

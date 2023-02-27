package laukaa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏奥拉赫蒂SuolahtiBarony struct {
	feud.BaseBarony
}

var BSuolahti苏奥拉赫蒂 feud.Barony = &苏奥拉赫蒂SuolahtiBarony{}

func init() {
    f := BSuolahti苏奥拉赫蒂.(*苏奥拉赫蒂SuolahtiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "suolahti",
		TitleName: "苏奥拉赫蒂",
		TitleCode: "b_suolahti",
	}
}

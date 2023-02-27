package kaisereia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多巴达DobadaBarony struct {
	feud.BaseBarony
}

var BDobada多巴达 feud.Barony = &多巴达DobadaBarony{}

func init() {
    f := BDobada多巴达.(*多巴达DobadaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dobada",
		TitleName: "多巴达",
		TitleCode: "b_dobada",
	}
}

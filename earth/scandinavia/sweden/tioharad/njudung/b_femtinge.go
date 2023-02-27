package njudung

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费姆廷厄FemtingeBarony struct {
	feud.BaseBarony
}

var BFemtinge费姆廷厄 feud.Barony = &费姆廷厄FemtingeBarony{}

func init() {
    f := BFemtinge费姆廷厄.(*费姆廷厄FemtingeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "femtinge",
		TitleName: "费姆廷厄",
		TitleCode: "b_femtinge",
	}
}

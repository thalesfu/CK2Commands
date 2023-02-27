package toropets

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰米扬斯克DemyanskBarony struct {
	feud.BaseBarony
}

var BDemyansk杰米扬斯克 feud.Barony = &杰米扬斯克DemyanskBarony{}

func init() {
    f := BDemyansk杰米扬斯克.(*杰米扬斯克DemyanskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "demyansk",
		TitleName: "杰米扬斯克",
		TitleCode: "b_demyansk",
	}
}

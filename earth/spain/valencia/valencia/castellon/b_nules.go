package castellon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 努莱斯NulesBarony struct {
	feud.BaseBarony
}

var BNules努莱斯 feud.Barony = &努莱斯NulesBarony{}

func init() {
    f := BNules努莱斯.(*努莱斯NulesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nules",
		TitleName: "努莱斯",
		TitleCode: "b_nules",
	}
}

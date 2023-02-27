package kalaus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰米诺DeminoBarony struct {
	feud.BaseBarony
}

var BDemino杰米诺 feud.Barony = &杰米诺DeminoBarony{}

func init() {
    f := BDemino杰米诺.(*杰米诺DeminoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "demino",
		TitleName: "杰米诺",
		TitleCode: "b_demino",
	}
}

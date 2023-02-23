package adrianopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德莫提卡DemotikaBarony struct {
	feud.BaseBarony
}

var BDemotika德莫提卡 feud.Barony = &德莫提卡DemotikaBarony{}

func init() {
	f := BDemotika德莫提卡.(*德莫提卡DemotikaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "demotika",
		TitleName: "德莫提卡",
		TitleCode: "b_demotika",
	}
}

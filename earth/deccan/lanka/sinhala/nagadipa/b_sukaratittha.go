package nagadipa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏迦罗咥他SukaratitthaBarony struct {
	feud.BaseBarony
}

var BSukaratittha苏迦罗咥他 feud.Barony = &苏迦罗咥他SukaratitthaBarony{}

func init() {
	f := BSukaratittha苏迦罗咥他.(*苏迦罗咥他SukaratitthaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sukaratittha",
		TitleName: "苏迦罗咥他",
		TitleCode: "b_sukaratittha",
	}
}

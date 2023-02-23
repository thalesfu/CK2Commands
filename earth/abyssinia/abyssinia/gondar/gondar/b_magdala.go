package gondar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马格达拉MagdalaBarony struct {
	feud.BaseBarony
}

var BMagdala马格达拉 feud.Barony = &马格达拉MagdalaBarony{}

func init() {
	f := BMagdala马格达拉.(*马格达拉MagdalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "magdala",
		TitleName: "马格达拉",
		TitleCode: "b_magdala",
	}
}

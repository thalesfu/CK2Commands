package maine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马耶讷MayenneBarony struct {
	feud.BaseBarony
}

var BMayenne马耶讷 feud.Barony = &马耶讷MayenneBarony{}

func init() {
	f := BMayenne马耶讷.(*马耶讷MayenneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mayenne",
		TitleName: "马耶讷",
		TitleCode: "b_mayenne",
	}
}

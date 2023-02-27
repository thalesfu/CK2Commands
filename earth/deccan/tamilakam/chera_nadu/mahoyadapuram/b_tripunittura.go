package mahoyadapuram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帝利补尼菟罗TripunitturaBarony struct {
	feud.BaseBarony
}

var BTripunittura帝利补尼菟罗 feud.Barony = &帝利补尼菟罗TripunitturaBarony{}

func init() {
    f := BTripunittura帝利补尼菟罗.(*帝利补尼菟罗TripunitturaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tripunittura",
		TitleName: "帝利补尼菟罗",
		TitleCode: "b_tripunittura",
	}
}

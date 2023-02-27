package french_leon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 朗代诺LanderneauBarony struct {
	feud.BaseBarony
}

var BLanderneau朗代诺 feud.Barony = &朗代诺LanderneauBarony{}

func init() {
    f := BLanderneau朗代诺.(*朗代诺LanderneauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "landerneau",
		TitleName: "朗代诺",
		TitleCode: "b_landerneau",
	}
}

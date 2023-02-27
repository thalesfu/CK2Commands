package kaliopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马蒂塔MadytaBarony struct {
	feud.BaseBarony
}

var BMadyta马蒂塔 feud.Barony = &马蒂塔MadytaBarony{}

func init() {
    f := BMadyta马蒂塔.(*马蒂塔MadytaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "madyta",
		TitleName: "马蒂塔",
		TitleCode: "b_madyta",
	}
}

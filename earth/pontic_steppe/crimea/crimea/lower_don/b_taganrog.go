package lower_don

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔甘罗格TaganrogBarony struct {
	feud.BaseBarony
}

var BTaganrog塔甘罗格 feud.Barony = &塔甘罗格TaganrogBarony{}

func init() {
    f := BTaganrog塔甘罗格.(*塔甘罗格TaganrogBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taganrog",
		TitleName: "塔甘罗格",
		TitleCode: "b_taganrog",
	}
}

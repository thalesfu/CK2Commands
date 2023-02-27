package tortosa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴勒米亚BalemiaBarony struct {
	feud.BaseBarony
}

var BBalemia巴勒米亚 feud.Barony = &巴勒米亚BalemiaBarony{}

func init() {
    f := BBalemia巴勒米亚.(*巴勒米亚BalemiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "balemia",
		TitleName: "巴勒米亚",
		TitleCode: "b_balemia",
	}
}

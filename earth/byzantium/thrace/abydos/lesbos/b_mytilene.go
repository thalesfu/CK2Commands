package lesbos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米蒂利尼MytileneBarony struct {
	feud.BaseBarony
}

var BMytilene米蒂利尼 feud.Barony = &米蒂利尼MytileneBarony{}

func init() {
	f := BMytilene米蒂利尼.(*米蒂利尼MytileneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mytilene",
		TitleName: "米蒂利尼",
		TitleCode: "b_mytilene",
	}
}

package molina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃尔卡萨尔ElcasarBarony struct {
	feud.BaseBarony
}

var BElcasar埃尔卡萨尔 feud.Barony = &埃尔卡萨尔ElcasarBarony{}

func init() {
	f := BElcasar埃尔卡萨尔.(*埃尔卡萨尔ElcasarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elcasar",
		TitleName: "埃尔卡萨尔",
		TitleCode: "b_elcasar",
	}
}

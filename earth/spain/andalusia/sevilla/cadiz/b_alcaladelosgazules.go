package cadiz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加苏莱斯堡AlcaladelosgazulesBarony struct {
	feud.BaseBarony
}

var BAlcaladelosgazules加苏莱斯堡 feud.Barony = &加苏莱斯堡AlcaladelosgazulesBarony{}

func init() {
	f := BAlcaladelosgazules加苏莱斯堡.(*加苏莱斯堡AlcaladelosgazulesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alcaladelosgazules",
		TitleName: "加苏莱斯堡",
		TitleCode: "b_alcaladelosgazules",
	}
}

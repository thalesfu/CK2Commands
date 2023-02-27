package orangallu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩兰格尔MolangoorBarony struct {
	feud.BaseBarony
}

var BMolangoor摩兰格尔 feud.Barony = &摩兰格尔MolangoorBarony{}

func init() {
    f := BMolangoor摩兰格尔.(*摩兰格尔MolangoorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "molangoor",
		TitleName: "摩兰格尔",
		TitleCode: "b_molangoor",
	}
}

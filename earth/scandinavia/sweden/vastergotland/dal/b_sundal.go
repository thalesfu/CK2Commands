package dal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 松达尔SundalBarony struct {
	feud.BaseBarony
}

var BSundal松达尔 feud.Barony = &松达尔SundalBarony{}

func init() {
	f := BSundal松达尔.(*松达尔SundalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sundal",
		TitleName: "松达尔",
		TitleCode: "b_sundal",
	}
}

package treviso

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥代尔佐OderzoBarony struct {
	feud.BaseBarony
}

var BOderzo奥代尔佐 feud.Barony = &奥代尔佐OderzoBarony{}

func init() {
	f := BOderzo奥代尔佐.(*奥代尔佐OderzoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oderzo",
		TitleName: "奥代尔佐",
		TitleCode: "b_oderzo",
	}
}

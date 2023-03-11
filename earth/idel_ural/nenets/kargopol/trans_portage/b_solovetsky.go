package trans_portage

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索洛韦茨基SolovetskyBarony struct {
	feud.BaseBarony
}

var BSolovetsky索洛韦茨基 feud.Barony = &索洛韦茨基SolovetskyBarony{}

func init() {
    f := BSolovetsky索洛韦茨基.(*索洛韦茨基SolovetskyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "solovetsky",
		TitleName: "索洛韦茨基",
		TitleCode: "b_solovetsky",
	}
}

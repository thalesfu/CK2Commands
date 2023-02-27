package asturias_de_oviedo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥维耶多OviedoBarony struct {
	feud.BaseBarony
}

var BOviedo奥维耶多 feud.Barony = &奥维耶多OviedoBarony{}

func init() {
    f := BOviedo奥维耶多.(*奥维耶多OviedoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oviedo",
		TitleName: "奥维耶多",
		TitleCode: "b_oviedo",
	}
}

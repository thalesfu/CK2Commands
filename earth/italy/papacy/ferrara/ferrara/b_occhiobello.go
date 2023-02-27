package ferrara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥基奥贝洛OcchiobelloBarony struct {
	feud.BaseBarony
}

var BOcchiobello奥基奥贝洛 feud.Barony = &奥基奥贝洛OcchiobelloBarony{}

func init() {
    f := BOcchiobello奥基奥贝洛.(*奥基奥贝洛OcchiobelloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "occhiobello",
		TitleName: "奥基奥贝洛",
		TitleCode: "b_occhiobello",
	}
}

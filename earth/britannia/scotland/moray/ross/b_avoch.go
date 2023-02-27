package ross

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥赫AvochBarony struct {
	feud.BaseBarony
}

var BAvoch奥赫 feud.Barony = &奥赫AvochBarony{}

func init() {
    f := BAvoch奥赫.(*奥赫AvochBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "avoch",
		TitleName: "奥赫",
		TitleCode: "b_avoch",
	}
}

package pest

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦茨VacBarony struct {
	feud.BaseBarony
}

var BVac瓦茨 feud.Barony = &瓦茨VacBarony{}

func init() {
    f := BVac瓦茨.(*瓦茨VacBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vac",
		TitleName: "瓦茨",
		TitleCode: "b_vac",
	}
}

package nagauda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙德拉MundwaBarony struct {
	feud.BaseBarony
}

var BMundwa蒙德拉 feud.Barony = &蒙德拉MundwaBarony{}

func init() {
	f := BMundwa蒙德拉.(*蒙德拉MundwaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mundwa",
		TitleName: "蒙德拉",
		TitleCode: "b_mundwa",
	}
}

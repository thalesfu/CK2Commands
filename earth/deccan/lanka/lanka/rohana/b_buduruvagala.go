package rohana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博杜瓦迦拉BuduruvagalaBarony struct {
	feud.BaseBarony
}

var BBuduruvagala博杜瓦迦拉 feud.Barony = &博杜瓦迦拉BuduruvagalaBarony{}

func init() {
    f := BBuduruvagala博杜瓦迦拉.(*博杜瓦迦拉BuduruvagalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "buduruvagala",
		TitleName: "博杜瓦迦拉",
		TitleCode: "b_buduruvagala",
	}
}

package grassland_cheremisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚尔夏勒YarcalliBarony struct {
	feud.BaseBarony
}

var BYarcalli亚尔夏勒 feud.Barony = &亚尔夏勒YarcalliBarony{}

func init() {
    f := BYarcalli亚尔夏勒.(*亚尔夏勒YarcalliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yarcalli",
		TitleName: "亚尔夏勒",
		TitleCode: "b_yarcalli",
	}
}

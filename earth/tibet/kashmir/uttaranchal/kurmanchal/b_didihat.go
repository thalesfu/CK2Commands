package kurmanchal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪迪哈特DidihatBarony struct {
	feud.BaseBarony
}

var BDidihat迪迪哈特 feud.Barony = &迪迪哈特DidihatBarony{}

func init() {
	f := BDidihat迪迪哈特.(*迪迪哈特DidihatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "didihat",
		TitleName: "迪迪哈特",
		TitleCode: "b_didihat",
	}
}

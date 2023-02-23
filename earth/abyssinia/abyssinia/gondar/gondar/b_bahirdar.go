package gondar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴赫达尔BahirdarBarony struct {
	feud.BaseBarony
}

var BBahirdar巴赫达尔 feud.Barony = &巴赫达尔BahirdarBarony{}

func init() {
	f := BBahirdar巴赫达尔.(*巴赫达尔BahirdarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bahirdar",
		TitleName: "巴赫达尔",
		TitleCode: "b_bahirdar",
	}
}

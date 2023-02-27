package banavasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉昆迪LakkundiBarony struct {
	feud.BaseBarony
}

var BLakkundi拉昆迪 feud.Barony = &拉昆迪LakkundiBarony{}

func init() {
    f := BLakkundi拉昆迪.(*拉昆迪LakkundiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lakkundi",
		TitleName: "拉昆迪",
		TitleCode: "b_lakkundi",
	}
}

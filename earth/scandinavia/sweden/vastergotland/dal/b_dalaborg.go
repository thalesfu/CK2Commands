package dal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达拉波雷DalaborgBarony struct {
	feud.BaseBarony
}

var BDalaborg达拉波雷 feud.Barony = &达拉波雷DalaborgBarony{}

func init() {
    f := BDalaborg达拉波雷.(*达拉波雷DalaborgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dalaborg",
		TitleName: "达拉波雷",
		TitleCode: "b_dalaborg",
	}
}

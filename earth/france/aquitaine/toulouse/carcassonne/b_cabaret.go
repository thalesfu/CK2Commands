package carcassonne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡巴雷CabaretBarony struct {
	feud.BaseBarony
}

var BCabaret卡巴雷 feud.Barony = &卡巴雷CabaretBarony{}

func init() {
    f := BCabaret卡巴雷.(*卡巴雷CabaretBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cabaret",
		TitleName: "卡巴雷",
		TitleCode: "b_cabaret",
	}
}

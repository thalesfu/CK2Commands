package trinkitat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 陶卡尔TokarBarony struct {
	feud.BaseBarony
}

var BTokar陶卡尔 feud.Barony = &陶卡尔TokarBarony{}

func init() {
    f := BTokar陶卡尔.(*陶卡尔TokarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tokar",
		TitleName: "陶卡尔",
		TitleCode: "b_tokar",
	}
}

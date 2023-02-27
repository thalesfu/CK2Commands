package messina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡塔拉蒂CatarattiBarony struct {
	feud.BaseBarony
}

var BCataratti卡塔拉蒂 feud.Barony = &卡塔拉蒂CatarattiBarony{}

func init() {
    f := BCataratti卡塔拉蒂.(*卡塔拉蒂CatarattiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cataratti",
		TitleName: "卡塔拉蒂",
		TitleCode: "b_cataratti",
	}
}

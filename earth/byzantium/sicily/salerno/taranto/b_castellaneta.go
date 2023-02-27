package taranto

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡斯泰拉内塔CastellanetaBarony struct {
	feud.BaseBarony
}

var BCastellaneta卡斯泰拉内塔 feud.Barony = &卡斯泰拉内塔CastellanetaBarony{}

func init() {
    f := BCastellaneta卡斯泰拉内塔.(*卡斯泰拉内塔CastellanetaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "castellaneta",
		TitleName: "卡斯泰拉内塔",
		TitleCode: "b_castellaneta",
	}
}

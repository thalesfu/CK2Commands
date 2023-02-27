package iasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃尔泽雷什蒂VarzarestiBarony struct {
	feud.BaseBarony
}

var BVarzaresti沃尔泽雷什蒂 feud.Barony = &沃尔泽雷什蒂VarzarestiBarony{}

func init() {
    f := BVarzaresti沃尔泽雷什蒂.(*沃尔泽雷什蒂VarzarestiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "varzaresti",
		TitleName: "沃尔泽雷什蒂",
		TitleCode: "b_varzaresti",
	}
}

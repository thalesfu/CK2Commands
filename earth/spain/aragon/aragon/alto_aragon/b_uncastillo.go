package alto_aragon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 温卡斯蒂略UncastilloBarony struct {
	feud.BaseBarony
}

var BUncastillo温卡斯蒂略 feud.Barony = &温卡斯蒂略UncastilloBarony{}

func init() {
    f := BUncastillo温卡斯蒂略.(*温卡斯蒂略UncastilloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uncastillo",
		TitleName: "温卡斯蒂略",
		TitleCode: "b_uncastillo",
	}
}

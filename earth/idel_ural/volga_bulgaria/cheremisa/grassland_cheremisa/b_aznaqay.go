package grassland_cheremisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿兹纳卡伊AznaqayBarony struct {
	feud.BaseBarony
}

var BAznaqay阿兹纳卡伊 feud.Barony = &阿兹纳卡伊AznaqayBarony{}

func init() {
    f := BAznaqay阿兹纳卡伊.(*阿兹纳卡伊AznaqayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aznaqay",
		TitleName: "阿兹纳卡伊",
		TitleCode: "b_aznaqay",
	}
}

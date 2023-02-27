package dhu_zabi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪拜DubaiBarony struct {
	feud.BaseBarony
}

var BDubai迪拜 feud.Barony = &迪拜DubaiBarony{}

func init() {
    f := BDubai迪拜.(*迪拜DubaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dubai",
		TitleName: "迪拜",
		TitleCode: "b_dubai",
	}
}

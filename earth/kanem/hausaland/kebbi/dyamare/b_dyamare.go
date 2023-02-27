package dyamare

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪亚马尔DyamareBarony struct {
	feud.BaseBarony
}

var BDyamare迪亚马尔 feud.Barony = &迪亚马尔DyamareBarony{}

func init() {
    f := BDyamare迪亚马尔.(*迪亚马尔DyamareBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dyamare",
		TitleName: "迪亚马尔",
		TitleCode: "b_dyamare",
	}
}

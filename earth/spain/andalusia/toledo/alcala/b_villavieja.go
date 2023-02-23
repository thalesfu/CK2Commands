package alcala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维拉韦利亚VillaviejaBarony struct {
	feud.BaseBarony
}

var BVillavieja维拉韦利亚 feud.Barony = &维拉韦利亚VillaviejaBarony{}

func init() {
	f := BVillavieja维拉韦利亚.(*维拉韦利亚VillaviejaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "villavieja",
		TitleName: "维拉韦利亚",
		TitleCode: "b_villavieja",
	}
}

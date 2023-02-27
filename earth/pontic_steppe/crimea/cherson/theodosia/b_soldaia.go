package theodosia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索尔得亚SoldaiaBarony struct {
	feud.BaseBarony
}

var BSoldaia索尔得亚 feud.Barony = &索尔得亚SoldaiaBarony{}

func init() {
    f := BSoldaia索尔得亚.(*索尔得亚SoldaiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "soldaia",
		TitleName: "索尔得亚",
		TitleCode: "b_soldaia",
	}
}

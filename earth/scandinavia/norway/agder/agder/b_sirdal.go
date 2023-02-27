package agder

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 西尔达尔SirdalBarony struct {
	feud.BaseBarony
}

var BSirdal西尔达尔 feud.Barony = &西尔达尔SirdalBarony{}

func init() {
    f := BSirdal西尔达尔.(*西尔达尔SirdalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sirdal",
		TitleName: "西尔达尔",
		TitleCode: "b_sirdal",
	}
}

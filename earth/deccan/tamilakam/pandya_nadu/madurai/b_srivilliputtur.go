package madurai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯里毗利布杜尔SrivilliputturBarony struct {
	feud.BaseBarony
}

var BSrivilliputtur斯里毗利布杜尔 feud.Barony = &斯里毗利布杜尔SrivilliputturBarony{}

func init() {
    f := BSrivilliputtur斯里毗利布杜尔.(*斯里毗利布杜尔SrivilliputturBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "srivilliputtur",
		TitleName: "斯里毗利布杜尔",
		TitleCode: "b_srivilliputtur",
	}
}

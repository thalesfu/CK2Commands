package syria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴鲁BalyBarony struct {
	feud.BaseBarony
}

var BBaly巴鲁 feud.Barony = &巴鲁BalyBarony{}

func init() {
    f := BBaly巴鲁.(*巴鲁BalyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baly",
		TitleName: "巴鲁",
		TitleCode: "b_baly",
	}
}

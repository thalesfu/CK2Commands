package sozopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 第纳尔DinarBarony struct {
	feud.BaseBarony
}

var BDinar第纳尔 feud.Barony = &第纳尔DinarBarony{}

func init() {
	f := BDinar第纳尔.(*第纳尔DinarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dinar",
		TitleName: "第纳尔",
		TitleCode: "b_dinar",
	}
}

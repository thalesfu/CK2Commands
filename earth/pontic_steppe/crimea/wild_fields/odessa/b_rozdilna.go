package odessa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉兹杰利纳亚RozdilnaBarony struct {
	feud.BaseBarony
}

var BRozdilna拉兹杰利纳亚 feud.Barony = &拉兹杰利纳亚RozdilnaBarony{}

func init() {
    f := BRozdilna拉兹杰利纳亚.(*拉兹杰利纳亚RozdilnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rozdilna",
		TitleName: "拉兹杰利纳亚",
		TitleCode: "b_rozdilna",
	}
}

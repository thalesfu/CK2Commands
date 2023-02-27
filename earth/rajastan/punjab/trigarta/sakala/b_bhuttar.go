package sakala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普多尔BhuttarBarony struct {
	feud.BaseBarony
}

var BBhuttar普多尔 feud.Barony = &普多尔BhuttarBarony{}

func init() {
    f := BBhuttar普多尔.(*普多尔BhuttarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhuttar",
		TitleName: "普多尔",
		TitleCode: "b_bhuttar",
	}
}

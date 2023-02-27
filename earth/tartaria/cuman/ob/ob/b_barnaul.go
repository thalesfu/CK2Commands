package ob

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴尔瑙尔BarnaulBarony struct {
	feud.BaseBarony
}

var BBarnaul巴尔瑙尔 feud.Barony = &巴尔瑙尔BarnaulBarony{}

func init() {
    f := BBarnaul巴尔瑙尔.(*巴尔瑙尔BarnaulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "barnaul",
		TitleName: "巴尔瑙尔",
		TitleCode: "b_barnaul",
	}
}

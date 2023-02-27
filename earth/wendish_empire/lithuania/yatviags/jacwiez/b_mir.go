package jacwiez

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 密尔MirBarony struct {
	feud.BaseBarony
}

var BMir密尔 feud.Barony = &密尔MirBarony{}

func init() {
    f := BMir密尔.(*密尔MirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mir",
		TitleName: "密尔",
		TitleCode: "b_mir",
	}
}

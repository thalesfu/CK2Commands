package verona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣马蒂诺SanmartinoBarony struct {
	feud.BaseBarony
}

var BSanmartino圣马蒂诺 feud.Barony = &圣马蒂诺SanmartinoBarony{}

func init() {
    f := BSanmartino圣马蒂诺.(*圣马蒂诺SanmartinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sanmartino",
		TitleName: "圣马蒂诺",
		TitleCode: "b_sanmartino",
	}
}

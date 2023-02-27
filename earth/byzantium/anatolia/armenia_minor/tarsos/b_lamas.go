package tarsos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉马斯LamasBarony struct {
	feud.BaseBarony
}

var BLamas拉马斯 feud.Barony = &拉马斯LamasBarony{}

func init() {
    f := BLamas拉马斯.(*拉马斯LamasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lamas",
		TitleName: "拉马斯",
		TitleCode: "b_lamas",
	}
}

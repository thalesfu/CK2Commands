package bourbon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 朗西LancyBarony struct {
	feud.BaseBarony
}

var BLancy朗西 feud.Barony = &朗西LancyBarony{}

func init() {
    f := BLancy朗西.(*朗西LancyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lancy",
		TitleName: "朗西",
		TitleCode: "b_lancy",
	}
}

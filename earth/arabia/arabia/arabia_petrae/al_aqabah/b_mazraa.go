package al_aqabah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迈兹拉阿MazraaBarony struct {
	feud.BaseBarony
}

var BMazraa迈兹拉阿 feud.Barony = &迈兹拉阿MazraaBarony{}

func init() {
    f := BMazraa迈兹拉阿.(*迈兹拉阿MazraaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mazraa",
		TitleName: "迈兹拉阿",
		TitleCode: "b_mazraa",
	}
}

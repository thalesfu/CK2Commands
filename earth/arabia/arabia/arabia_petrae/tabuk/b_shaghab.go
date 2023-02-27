package tabuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舍盖卜ShaghabBarony struct {
	feud.BaseBarony
}

var BShaghab舍盖卜 feud.Barony = &舍盖卜ShaghabBarony{}

func init() {
    f := BShaghab舍盖卜.(*舍盖卜ShaghabBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shaghab",
		TitleName: "舍盖卜",
		TitleCode: "b_shaghab",
	}
}

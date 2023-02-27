package narke

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 厄勒布鲁OrebroBarony struct {
	feud.BaseBarony
}

var BOrebro厄勒布鲁 feud.Barony = &厄勒布鲁OrebroBarony{}

func init() {
    f := BOrebro厄勒布鲁.(*厄勒布鲁OrebroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "orebro",
		TitleName: "厄勒布鲁",
		TitleCode: "b_orebro",
	}
}

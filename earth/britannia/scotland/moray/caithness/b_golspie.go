package caithness

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈尔斯皮GolspieBarony struct {
	feud.BaseBarony
}

var BGolspie戈尔斯皮 feud.Barony = &戈尔斯皮GolspieBarony{}

func init() {
    f := BGolspie戈尔斯皮.(*戈尔斯皮GolspieBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "golspie",
		TitleName: "戈尔斯皮",
		TitleCode: "b_golspie",
	}
}

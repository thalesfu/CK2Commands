package sutai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达尔维Darvi_sutaiBarony struct {
	feud.BaseBarony
}

var BDarvi_sutai达尔维 feud.Barony = &达尔维Darvi_sutaiBarony{}

func init() {
    f := BDarvi_sutai达尔维.(*达尔维Darvi_sutaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "darvi_sutai",
		TitleName: "达尔维",
		TitleCode: "b_darvi_sutai",
	}
}

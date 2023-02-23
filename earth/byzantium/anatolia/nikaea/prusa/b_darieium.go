package prusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达里埃翁DarieiumBarony struct {
	feud.BaseBarony
}

var BDarieium达里埃翁 feud.Barony = &达里埃翁DarieiumBarony{}

func init() {
	f := BDarieium达里埃翁.(*达里埃翁DarieiumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "darieium",
		TitleName: "达里埃翁",
		TitleCode: "b_darieium",
	}
}

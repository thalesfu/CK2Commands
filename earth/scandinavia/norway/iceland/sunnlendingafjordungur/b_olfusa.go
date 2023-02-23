package sunnlendingafjordungur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 厄尔菲绍OlfusaBarony struct {
	feud.BaseBarony
}

var BOlfusa厄尔菲绍 feud.Barony = &厄尔菲绍OlfusaBarony{}

func init() {
	f := BOlfusa厄尔菲绍.(*厄尔菲绍OlfusaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "olfusa",
		TitleName: "厄尔菲绍",
		TitleCode: "b_olfusa",
	}
}

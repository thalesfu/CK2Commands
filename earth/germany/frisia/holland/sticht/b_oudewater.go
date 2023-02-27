package sticht

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥德瓦特OudewaterBarony struct {
	feud.BaseBarony
}

var BOudewater奥德瓦特 feud.Barony = &奥德瓦特OudewaterBarony{}

func init() {
    f := BOudewater奥德瓦特.(*奥德瓦特OudewaterBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oudewater",
		TitleName: "奥德瓦特",
		TitleCode: "b_oudewater",
	}
}

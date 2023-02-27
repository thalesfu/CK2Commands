package sacz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮尔兹诺PilznoBarony struct {
	feud.BaseBarony
}

var BPilzno皮尔兹诺 feud.Barony = &皮尔兹诺PilznoBarony{}

func init() {
    f := BPilzno皮尔兹诺.(*皮尔兹诺PilznoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pilzno",
		TitleName: "皮尔兹诺",
		TitleCode: "b_pilzno",
	}
}

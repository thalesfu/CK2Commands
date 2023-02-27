package qalqut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波罗波那迦迪ParappanangadiBarony struct {
	feud.BaseBarony
}

var BParappanangadi波罗波那迦迪 feud.Barony = &波罗波那迦迪ParappanangadiBarony{}

func init() {
    f := BParappanangadi波罗波那迦迪.(*波罗波那迦迪ParappanangadiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "parappanangadi",
		TitleName: "波罗波那迦迪",
		TitleCode: "b_parappanangadi",
	}
}

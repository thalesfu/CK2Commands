package la_marche

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布尔加讷夫BourganeufBarony struct {
	feud.BaseBarony
}

var BBourganeuf布尔加讷夫 feud.Barony = &布尔加讷夫BourganeufBarony{}

func init() {
    f := BBourganeuf布尔加讷夫.(*布尔加讷夫BourganeufBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bourganeuf",
		TitleName: "布尔加讷夫",
		TitleCode: "b_bourganeuf",
	}
}

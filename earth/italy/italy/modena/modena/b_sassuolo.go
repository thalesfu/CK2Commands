package modena

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨索洛SassuoloBarony struct {
	feud.BaseBarony
}

var BSassuolo萨索洛 feud.Barony = &萨索洛SassuoloBarony{}

func init() {
    f := BSassuolo萨索洛.(*萨索洛SassuoloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sassuolo",
		TitleName: "萨索洛",
		TitleCode: "b_sassuolo",
	}
}

package safed

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 采法特SafedBarony struct {
	feud.BaseBarony
}

var BSafed采法特 feud.Barony = &采法特SafedBarony{}

func init() {
    f := BSafed采法特.(*采法特SafedBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "safed",
		TitleName: "采法特",
		TitleCode: "b_safed",
	}
}

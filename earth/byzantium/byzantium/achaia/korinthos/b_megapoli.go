package korinthos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巨城MegapoliBarony struct {
	feud.BaseBarony
}

var BMegapoli巨城 feud.Barony = &巨城MegapoliBarony{}

func init() {
	f := BMegapoli巨城.(*巨城MegapoliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "megapoli",
		TitleName: "巨城",
		TitleCode: "b_megapoli",
	}
}

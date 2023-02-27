package gizeh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利什特EllishtBarony struct {
	feud.BaseBarony
}

var BEllisht利什特 feud.Barony = &利什特EllishtBarony{}

func init() {
    f := BEllisht利什特.(*利什特EllishtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ellisht",
		TitleName: "利什特",
		TitleCode: "b_ellisht",
	}
}

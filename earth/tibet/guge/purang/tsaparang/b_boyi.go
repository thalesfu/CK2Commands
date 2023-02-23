package tsaparang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博依BoyiBarony struct {
	feud.BaseBarony
}

var BBoyi博依 feud.Barony = &博依BoyiBarony{}

func init() {
	f := BBoyi博依.(*博依BoyiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "boyi",
		TitleName: "博依",
		TitleCode: "b_boyi",
	}
}

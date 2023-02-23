package berg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴门BarmenBarony struct {
	feud.BaseBarony
}

var BBarmen巴门 feud.Barony = &巴门BarmenBarony{}

func init() {
	f := BBarmen巴门.(*巴门BarmenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "barmen",
		TitleName: "巴门",
		TitleCode: "b_barmen",
	}
}

package benevento

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣安杰洛SanangeloBarony struct {
	feud.BaseBarony
}

var BSanangelo圣安杰洛 feud.Barony = &圣安杰洛SanangeloBarony{}

func init() {
	f := BSanangelo圣安杰洛.(*圣安杰洛SanangeloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sanangelo",
		TitleName: "圣安杰洛",
		TitleCode: "b_sanangelo",
	}
}
